package main

import (
	"fmt"
	"github.com/jamesmoriarty/gomem"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/MaxwelMazur/csboost/internal/repository"
	goxymemmory "github.com/aditkumar1/goxymemory"
)

const (
	enableVal  = 1
	disableVal = 0

	VK_SHIFT         = 0x10 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	VK_CONTROL       = 0x11
	CSGO_FORCEATTACK = 0x6

	VK_SPACE         = 0x20   // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	CSGO_FL_ONGROUND = 1 << 0 // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
	CSGO_FORCEJUMP   = 0x6    // https://github.com/ValveSoftware/source-sdk-2013/blob/0d8dceea4310fde5706b3ce1c70609d72a38efdf/sp/src/game/shared/sdk/sdk_playeranimstate.cpp#L517

)

func main() {
	println("starting...")
	offsets, err := repository.GetNewOffset()
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

	// [ radar ]
	go func() {
		for {
			clientAddress, _ := dm.GetModuleFromName("client.dll")
			for i := 1; i < 32; i++ {
				entity, _ := dm.Read((uint)(clientAddress)+(offsets.Signatures.DwEntityList+(uint)(i*0x10)), goxymemmory.UINT)
				if entity.Value.(uint32) > 0 {
					entityTeamId, _ := dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MITeamNum, goxymemmory.UINT)
					if entityTeamId.Value.(uint32) == 2 {
						dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)), goxymemmory.Data{Value: enableVal, DataType: goxymemmory.UINT})
					}
					if entityTeamId.Value.(uint32) == 3 {
						dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)), goxymemmory.Data{Value: enableVal, DataType: goxymemmory.UINT})
					}
				}
			}
		}
	}()

	// [ trigger ]
	go func() {
		for {
			rtNum := reactionTime()
			fmt.Printf("\033[2K\rTempo gerado aleatoriamente: %d milisegundos", rtNum)
			time.Sleep(time.Duration(rtNum) * time.Millisecond)
			clientAddress, _ := dm.GetModuleFromName("client.dll")
			if gomem.IsKeyDown(VK_SHIFT) || gomem.IsKeyDown(VK_CONTROL) {
				localPlayer, _ := dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), goxymemmory.UINT)
				entity, _ := dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MICrosshairId)), goxymemmory.UINT)
				if entity.Value.(uint32) > 0 && entity.Value.(uint32) <= 64 {
					dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceAttack)), goxymemmory.Data{Value: CSGO_FORCEATTACK, DataType: goxymemmory.INT})
				}
			}
		}
	}()

	// [ bhop ]
	go func() {
		for {
			clientAddress, _ := dm.GetModuleFromName("client.dll")
			if gomem.IsKeyDown(VK_SPACE) {
				localPlayer, _ := dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), goxymemmory.UINT)
				flags, _ := dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MFFlags)), goxymemmory.UINT)
				if (flags.Value.(uint32) & CSGO_FL_ONGROUND) > 0 {
					dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceJump)), goxymemmory.Data{Value: CSGO_FORCEJUMP, DataType: goxymemmory.INT})
				}
			}
		}
	}()

	select {}
}

func reactionTime() int {
	num := rand.Intn(84) + 133
	if float64(num) < 0.7*float64(171-133)+133 {
		num += rand.Intn(47)
	}
	return num
}
