package model

type Cacher interface {
	Get(string) ([]byte, error)
	Put(string, []byte) error
}
