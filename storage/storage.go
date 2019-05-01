package storage

import (
	"fmt"

	"github.com/clintjedwards/cursor/config"
)

// EngineType type represents the different possible storage engines available
type EngineType string

const (
	// StorageEngineBoltDB represents a boltDB storage engine.
	// A file based key-value store.(https://github.com/boltdb/bolt)
	StorageEngineBoltDB EngineType = "boltdb"
)

// Engine represents backend storage implementations where items can be persisted
type Engine interface {
	Init(config *config.Config) error
	GetAll() map[string][]byte
	Get(key string) ([]byte, error)
	Add(value []byte) error
	Update(key string, newValue []byte) error
	Delete(key string) error
}

// InitStorage creates a storage object with the appropriate engine
func InitStorage(engineType EngineType) (Engine, error) {

	switch engineType {
	case StorageEngineBoltDB:
		config, err := config.FromEnv()
		if err != nil {
			return nil, err
		}

		boltDBStorageEngine := boltDB{}
		err = boltDBStorageEngine.Init(config)
		if err != nil {
			return nil, err
		}

		return &boltDBStorageEngine, nil
	default:
		return nil, fmt.Errorf("storage backend not implemented: %s", engineType)
	}
}