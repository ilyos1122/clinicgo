package models

type CreateClinic struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
	City string `json:"city"`
}

type Clinic struct {
	ClinicID string `json:"clinic_id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	City     string `json:"city"`
}
