package controller

import (
	"log"
	"simple-clinic/models"
)

func (c *conn) CreateClinic(req *models.CreateClinic) (*models.Clinic, error) {

	log.Printf("Create clinin resp >>>>> %+v\n", req)
	resp, err := c.storage.Clinic().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
