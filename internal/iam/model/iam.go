package model

import (
	"github.com/caos/zitadel/internal/domain"
	es_models "github.com/caos/zitadel/internal/eventstore/v1/models"
)

type Step int

const (
	Step1 Step = iota + 1
	Step2
	Step3
	Step4
	Step5
	Step6
	Step7
	Step8
	Step9
	Step10
	//StepCount marks the the length of possible steps (StepCount-1 == last possible step)
	StepCount
)

type IAM struct {
	es_models.ObjectRoot
	GlobalOrgID                     string
	IAMProjectID                    string
	SetUpDone                       domain.Step
	SetUpStarted                    domain.Step
	Members                         []*IAMMember
	IDPs                            []*IDPConfig
	DefaultLoginPolicy              *LoginPolicy
	DefaultLabelPolicy              *LabelPolicy
	DefaultOrgIAMPolicy             *OrgIAMPolicy
	DefaultPasswordComplexityPolicy *PasswordComplexityPolicy
	DefaultPasswordAgePolicy        *PasswordAgePolicy
	DefaultLockoutPolicy            *LockoutPolicy
	DefaultMailTemplate             *MailTemplate
	DefaultMailTexts                []*MailText
}

func (iam *IAM) GetMember(userID string) (int, *IAMMember) {
	for i, m := range iam.Members {
		if m.UserID == userID {
			return i, m
		}
	}
	return -1, nil
}

func (iam *IAM) GetIDP(idpID string) (int, *IDPConfig) {
	for i, idp := range iam.IDPs {
		if idp.IDPConfigID == idpID {
			return i, idp
		}
	}
	return -1, nil
}

func (iam *IAM) GetDefaultMailText(mailTextType string, language string) (int, *MailText) {
	for i, m := range iam.DefaultMailTexts {
		if m.MailTextType == mailTextType && m.Language == language {
			return i, m
		}
	}
	return -1, nil
}
