package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/jamesmoriarty/gomem"
	"github.com/maxwelbm/gorwmem"
	"math/rand"
	"time"
)

func Trigger(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		rtNum := reactionTime()
		time.Sleep(time.Duration(rtNum) * time.Millisecond)
		clientAddress, err := dm.GetModuleFromName("client.dll")
		if err != nil {
			fmt.Printf("Failed reading module client.dll. %s", err)
		}
		if gomem.IsKeyDown(vkShift) || gomem.IsKeyDown(vkControl) {
			var localPlayer gorwmem.Data
			localPlayer, err = dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), gorwmem.UINT)
			if err != nil {
				fmt.Printf("Failed reading memory localPlayer. %s", err)
			}
			var entity gorwmem.Data
			entity, err = dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MICrosshairId)), gorwmem.UINT)
			if err != nil {
				fmt.Printf("Failed reading memory entity. %s", err)
			}
			if entity.Value.(uint32) > 0 && entity.Value.(uint32) <= 64 {
				if err = dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceAttack)),
					gorwmem.Data{Value: csgoForceAttack, DataType: gorwmem.INT}); err != nil {
					fmt.Printf("Failed writing memory. %s", err)
					continue
				}
			}
		}
	}
}

func reactionTime() int {
	num := rand.Intn(84) + 133
	if float64(num) < 0.7*float64(171-133)+133 {
		num += rand.Intn(47)
	}
	return num
}
