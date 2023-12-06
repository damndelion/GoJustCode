package main

import (
	"awesomeProject/Lecture11/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Println(cfg.App.Name)    //returns go-clean-template
	fmt.Println(cfg.App.Version) // 1.0.0
	fmt.Println(cfg.HTTP.Port)   // 8080
	fmt.Println(cfg.Log.Level)   // debug
	fmt.Println(cfg.PG.URL)      // postgres://postgres:postgres@localhost:5432/postgres
}
