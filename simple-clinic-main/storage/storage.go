package storage

import "simple-clinic/models"

type StorageI interface {
	Clinic() ClinicRepoI
}

type ClinicRepoI interface {
	Create(req *models.CreateClinic) (*models.Clinic, error)
}
