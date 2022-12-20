package secret_manager

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"reservation-api/pkg/env"
)

var (
	mountPath  = "secret"
	secretPath = "reservation-api-secrets"
)

type SecretManager struct {
	client *vault.Client
}

func New() *SecretManager {

	config := vault.DefaultConfig()
	config.Address = env.GetFromDotENV("vault_addr")
	client, err := vault.NewClient(config)

	if err != nil {
		panic(err.Error())
	}

	client.SetToken(env.GetFromDotENV("vault_token"))
	return &SecretManager{
		client: client,
	}
}

func (s *SecretManager) Put(ctx context.Context, key, val string) error {

	secretData := make(map[string]interface{})
	secretData[key] = val

	_, err := s.client.KVv2(mountPath).Put(ctx, secretPath, secretData)
	return err
}

func (s *SecretManager) Get(ctx context.Context, key string) (interface{}, error) {

	secret, err := s.client.KVv2(mountPath).Get(ctx, secretPath)
	if err != nil {
		return nil, err
	}

	value, ok := secret.Data[key].(string)
	if !ok {
		return nil, fmt.Errorf(
			"value type assertion failed: %T %#v",
			secret.Data[key],
			secret.Data[key],
		)
	}

	return value, nil

}
