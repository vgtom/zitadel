package configuration

import (
	"time"

	"github.com/caos/orbos/mntr"
	"github.com/caos/orbos/pkg/kubernetes"
	"github.com/caos/orbos/pkg/kubernetes/resources/configmap"
	"github.com/caos/orbos/pkg/kubernetes/resources/secret"
	"github.com/caos/orbos/pkg/labels"
	"github.com/caos/zitadel/operator"
	"github.com/caos/zitadel/operator/kinds/iam/zitadel/configuration/users"
	"github.com/caos/zitadel/operator/kinds/iam/zitadel/database"
)

type ConsoleEnv struct {
	AuthServiceURL string `json:"authServiceUrl"`
	MgmtServiceURL string `json:"mgmtServiceUrl"`
	Issuer         string `json:"issuer"`
	ClientID       string `json:"clientid"`
}

const (
	googleServiceAccountJSONPath               = "google-serviceaccount-key.json"
	zitadelKeysPath                            = "zitadel-keys.yaml"
	timeout                      time.Duration = 60
)

func AdaptFunc(
	monitor mntr.Monitor,
	componentLabels *labels.Component,
	namespace string,
	desired *Configuration,
	cmName string,
	certPath string,
	secretName string,
	secretPath string,
	consoleCMName string,
	secretVarsName string,
	secretPasswordName string,
	necessaryUsers map[string]string,
	getClientID func() string,
	dbClient database.ClientInt,
) (
	operator.QueryFunc,
	operator.DestroyFunc,
	func(k8sClient kubernetes.ClientInt, queried map[string]interface{}) map[string]string,
	error,
) {
	internalMonitor := monitor.WithField("component", "configuration")

	literalsSecret := literalsSecret(desired, googleServiceAccountJSONPath, zitadelKeysPath)
	literalsSecretVars := literalsSecretVars(desired)

	destroyCM, err := configmap.AdaptFuncToDestroy(namespace, cmName)
	if err != nil {
		return nil, nil, nil, err
	}
	destroyS, err := secret.AdaptFuncToDestroy(namespace, secretName)
	if err != nil {
		return nil, nil, nil, err
	}
	destroyCCM, err := configmap.AdaptFuncToDestroy(namespace, consoleCMName)
	if err != nil {
		return nil, nil, nil, err
	}
	destroySV, err := secret.AdaptFuncToDestroy(namespace, secretVarsName)
	if err != nil {
		return nil, nil, nil, err
	}
	destroySP, err := secret.AdaptFuncToDestroy(namespace, secretPasswordName)
	if err != nil {
		return nil, nil, nil, err
	}

	_, destroyUser, err := users.AdaptFunc(internalMonitor, necessaryUsers, dbClient)
	if err != nil {
		return nil, nil, nil, err
	}

	destroyers := []operator.DestroyFunc{
		destroyUser,
		operator.ResourceDestroyToZitadelDestroy(destroyS),
		operator.ResourceDestroyToZitadelDestroy(destroyCM),
		operator.ResourceDestroyToZitadelDestroy(destroyCCM),
		operator.ResourceDestroyToZitadelDestroy(destroySV),
		operator.ResourceDestroyToZitadelDestroy(destroySP),
	}

	return func(k8sClient kubernetes.ClientInt, queried map[string]interface{}) (operator.EnsureFunc, error) {
			queryUser, _, err := users.AdaptFunc(internalMonitor, necessaryUsers, dbClient)
			if err != nil {
				return nil, err
			}
			queryS, err := secret.AdaptFuncToEnsure(namespace, labels.MustForName(componentLabels, secretName), literalsSecret)
			if err != nil {
				return nil, err
			}
			querySV, err := secret.AdaptFuncToEnsure(namespace, labels.MustForName(componentLabels, secretVarsName), literalsSecretVars)
			if err != nil {
				return nil, err
			}
			querySP, err := secret.AdaptFuncToEnsure(namespace, labels.MustForName(componentLabels, secretPasswordName), necessaryUsers)
			if err != nil {
				return nil, err
			}

			queryCCM, err := configmap.AdaptFuncToEnsure(
				namespace,
				consoleCMName,
				labels.MustForNameK8SMap(componentLabels, consoleCMName),
				literalsConsoleCM(
					getClientID(),
					desired.DNS,
					k8sClient,
					namespace,
					consoleCMName,
				),
			)
			if err != nil {
				return nil, err
			}

			queryCM, err := configmap.AdaptFuncToEnsure(
				namespace,
				cmName,
				labels.MustForNameK8SMap(componentLabels, cmName),
				literalsConfigMap(
					desired,
					necessaryUsers,
					certPath,
					secretPath,
					googleServiceAccountJSONPath,
					zitadelKeysPath,
					queried,
				),
			)
			if err != nil {
				return nil, err
			}

			queriers := []operator.QueryFunc{
				queryUser,
				operator.ResourceQueryToZitadelQuery(queryS),
				operator.ResourceQueryToZitadelQuery(queryCCM),
				operator.ResourceQueryToZitadelQuery(querySV),
				operator.ResourceQueryToZitadelQuery(querySP),
				operator.ResourceQueryToZitadelQuery(queryCM),
			}

			return operator.QueriersToEnsureFunc(internalMonitor, false, queriers, k8sClient, queried)
		},
		operator.DestroyersToDestroyFunc(internalMonitor, destroyers),
		func(k8sClient kubernetes.ClientInt, queried map[string]interface{}) map[string]string {
			return map[string]string{
				secretName:         getHash(literalsSecret),
				secretVarsName:     getHash(literalsSecretVars),
				secretPasswordName: getHash(necessaryUsers),
				cmName: getHash(
					literalsConfigMap(
						desired,
						necessaryUsers,
						certPath,
						secretPath,
						googleServiceAccountJSONPath,
						zitadelKeysPath,
						queried,
					),
				),
				consoleCMName: getHash(
					literalsConsoleCM(
						getClientID(),
						desired.DNS,
						k8sClient,
						namespace,
						consoleCMName,
					),
				),
			}
		},
		nil
}