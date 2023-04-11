package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	goxymemmory "github.com/aditkumar1/goxymemory"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	OffsetsURL = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.json"
	enableVal  = 1
	disableVal = 0
)

func GetNewOffset() (*model.OffSet, error) {
	resp, err := http.Get(OffsetsURL)
	if err != nil {
		return &model.OffSet{}, fmt.Errorf("fail to get offset. Error - %v. Using default offsets. Cheat may not work", err)
	}
	defer resp.Body.Close()

	strBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &model.OffSet{}, fmt.Errorf("unable to parse returned offset json. Error - %v. Using default offsets. Cheat may not work", err)
	}

	var offSet model.OffSet
	err = json.Unmarshal(strBytes, &offSet)
	if err != nil {
		return &model.OffSet{}, fmt.Errorf("unable to parse returned offset json. Error - %v. Using default offsets. Cheat may not work", err)
	}
	return &offSet, nil
}

func main() {
	println("starting...")
	offsets, err := GetNewOffset()
	if err != nil {
		log.Println("problemas in get offsets...")
	}

	// panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Encountered error - %v , ensure CS GO application running before starting this application.\n", r)
			log.Println("Exiting application")
			os.Exit(1)
		}
	}()

	log.Printf("launching cheat ......")
	dm := goxymemmory.DataManager("csgo.exe")

	log.Printf("offset updated ......")
	log.Printf("running ......")

	// hack loop
	for {
		clientAddress, _ := dm.GetModuleFromName("client.dll")
		for i := 1; i < 32; i++ {
			entity, _ := dm.Read((uint)(clientAddress)+(offsets.Signatures.DwEntityList+(uint)(i*0x10)), goxymemmory.UINT)
			if entity.Value.(uint32) > 0 {
				entityTeamId, _ := dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MITeamNum, goxymemmory.UINT)
				if entityTeamId.Value.(uint32) == 3 {
					dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)), goxymemmory.Data{Value: enableVal, DataType: goxymemmory.UINT})
				}
			}
		}
	}
}
