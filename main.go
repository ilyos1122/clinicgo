package main

import (
	"fmt"
	"log"
	"simple-clinic/config"
	"simple-clinic/controller"
	"simple-clinic/models"
	"simple-clinic/storage/jsondb"
)

func main() {

	var cfg = config.Load()

	jsondb, err := jsondb.NewJsonDbConnection(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, jsondb)

	resp, err := con.CreateClinic(
		&models.CreateClinic{
			Name: "Akfa clinic",
			Logo: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.vectorstock.com%2Froyalty-free-vector%2Fclinic-care-logo-icon-design-vector-22560856&psig=AOvVaw0rmUsoAItEKyY5eJKams9h&ust=1696684557507000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCIiQ1NTA4YEDFQAAAAAdAAAAABAE",
			City: "Tashkent",
		},
	)
	if err != nil {
		log.Println("Create clinic:", err.Error())
		return
	}

	fmt.Println(resp)
}
