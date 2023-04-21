package main

import (
	"github.com/MaxwelMazur/csboost/internal/process"
	"github.com/MaxwelMazur/csboost/internal/repository"
	"github.com/maxwelbm/gorwmem"
	"log"
	"math/rand"
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

	log.Printf("starting ...")
	log.Printf("launching cheat ......")
	log.Printf("offset updated ......")
	log.Printf("running ......")

	go process.Radar(dm, offsets)
	go process.Wall(dm, offsets)
	go process.BHop(dm, offsets)
	go process.Trigger(dm, offsets)

	select {}
}

func reactionTime() int {
	num := rand.Intn(84) + 133
	if float64(num) < 0.7*float64(171-133)+133 {
		num += rand.Intn(47)
	}
	return num
}
