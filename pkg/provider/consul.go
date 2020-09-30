package provider

type ConsulExecutor interface {
	KVGet(key string, v interface{}) error
	KVList(prefix string) (map[string][]byte, error)
}