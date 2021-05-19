package domain

type Key struct{}

type KeyProvider interface {
	GetKeys() ([]Key, error)
}
