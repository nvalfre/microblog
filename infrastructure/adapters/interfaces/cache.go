package interfaces

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration int) error
	Delete(key string) error
}
