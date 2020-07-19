package parameter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/rs/zerolog/log"
)

type IParameterStore interface {
	Load(string) (string, error)
}

type SSMParameterStore struct{}

func NewSSMParameterStore() SSMParameterStore {
	return SSMParameterStore{}
}

func (self SSMParameterStore) Load(path string) (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}

	ssmsvc := ssm.New(sess)
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}

	value := *param.Parameter.Value

	log.Printf("SSM parameter downloaded: %s", path)

	return value, nil
}
