package jsondb

import (
	"os"
	"simple-clinic/config"
	"simple-clinic/storage"
)

type Store struct {
	clinic storage.ClinicRepoI
}

func NewJsonDbConnection(cfg *config.Config) (storage.StorageI, error) {

	clinicFile, err := os.Open(cfg.ClinicPath)
	if err != nil {
		return nil, err
	}

	return &Store{
		clinic: NewClinicRepo(cfg, clinicFile),
	}, nil
}

func (s Store) Clinic() storage.ClinicRepoI {
	return s.clinic
}
