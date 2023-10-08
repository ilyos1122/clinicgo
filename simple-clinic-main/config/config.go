package config

type Config struct {
	ClinicPath string
}

func Load() Config {

	var cfg Config

	cfg.ClinicPath = "./data/clinic.json"

	return cfg
}
