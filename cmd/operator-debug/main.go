package main

import (
	"flag"
	"io/ioutil"

	"github.com/caos/orbos/mntr"
	"github.com/caos/orbos/pkg/kubernetes"
	"github.com/caos/zitadel/operator/helpers"
	"github.com/caos/zitadel/operator/start"
)

func main() {
	orbconfig := flag.String("orbconfig", "~/.orb/config", "The orbconfig file to use")
	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "The kubeconfig file to use")
	version := flag.String("version", "", "The ZITADEL version to deploy")
	verbose := flag.Bool("verbose", false, "Print debug levelled logs")

	flag.Parse()

	monitor := mntr.Monitor{
		OnInfo:   mntr.LogMessage,
		OnChange: mntr.LogMessage,
		OnError:  mntr.LogError,
	}

	if *version == "" {
		panic("flag --version is missing")
	}

	if *verbose {
		monitor = monitor.Verbose()
	}

	kc, err := ioutil.ReadFile(helpers.PruneHome(*kubeconfig))
	if err != nil {
		panic(err)
	}

	if err := start.Operator(
		monitor,
		helpers.PruneHome(*orbconfig),
		kubernetes.NewK8sClient(monitor, strPtr(string(kc))),
		"./migrations/cockroach/",
		version,
	); err != nil {
		panic(err)
	}
}

func strPtr(str string) *string {
	return &str
}