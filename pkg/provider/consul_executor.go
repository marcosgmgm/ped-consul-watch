package provider

import (
	"encoding/json"
	consulApi "github.com/hashicorp/consul/api"
)

type consulExecutor struct {
	clientApi *consulApi.Client
}

func NewConsulExecutor() (ConsulExecutor, error) {
	consulApi, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		return nil, err
	}
	ce := consulExecutor{
		clientApi: consulApi,
	}
	return ce, nil
}

func (ce consulExecutor) KVGet(key string, v interface{}) error {
	qo := &consulApi.QueryOptions{}
	kp, _, err := ce.clientApi.KV().Get(key, qo)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(kp.Value, v); err != nil {
		return err
	}
	return nil
}

func (ce consulExecutor) KVList(prefix string) (map[string][]byte, error) {
	qo := &consulApi.QueryOptions{}
	kp, _, err := ce.clientApi.KV().List(prefix, qo)
	if err != nil {
		return nil, err
	}
	response := make(map[string][]byte)
	for _, k := range kp {
		response[k.Key] = k.Value
	}
	return response, nil
}

