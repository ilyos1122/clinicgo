package jsondb

import (
	"encoding/json"
	"io"
	"os"

	"simple-clinic/config"
	"simple-clinic/models"
	"simple-clinic/storage"

	"github.com/google/uuid"
)

type clinicRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewClinicRepo(cfg *config.Config, file *os.File) storage.ClinicRepoI {
	return &clinicRepo{cfg: cfg, file: file}
}

func (c *clinicRepo) Create(req *models.CreateClinic) (*models.Clinic, error) {

	body, err := io.ReadAll(c.file)
	if err != nil {
		return nil, err
	}

	var clinics []*models.Clinic

	err = json.Unmarshal(body, &clinics)
	if err != nil {
		return nil, err
	}

	var resp = &models.Clinic{
		ClinicID: uuid.New().String(),
		Name:     req.Name,
		Logo:     req.Logo,
		City:     req.City,
	}

	clinics = append(clinics, resp)

	body, err = json.MarshalIndent(clinics, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(c.cfg.ClinicPath, body, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
