package main

import (
	"github.com/MaxwelMazur/csboost/internal/process"
	"github.com/MaxwelMazur/csboost/internal/repository"
	"github.com/maxwelbm/gorwmem"
	"log"
	"os"
)

func main() {
	log.Printf("Starting csboost ...")
	offsets, err := repository.GetOffset()
	if err != nil {
		log.Println("problemas in get offsets ...")
		log.Println("Exiting application")
		os.Exit(1)
	}

	dm := gorwmem.GetDataManager("csgo.exe")
	if !dm.IsOpen {
		log.Printf("CS GO application running before starting this application.\n")
		log.Println("Exiting application")
		os.Exit(1)
	}

	log.Printf("Running process ...")
	go process.Radar(dm, offsets)
	go process.Wall(dm, offsets)
	go process.BHop(dm, offsets)
	go process.Trigger(dm, offsets)

	select {}
}
