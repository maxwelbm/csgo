package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/maxwelbm/gorwmem"
)

func BHop(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, err := dm.GetModuleFromName("client.dll")
		if err != nil {
			fmt.Printf("Failed reading module client.dll. %s", err)
		}
		if gorwmem.IsKeyDown(vkSpace) {
			var localPlayer gorwmem.Data
			localPlayer, err = dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), 0, gorwmem.UINT)
			if err != nil {
				fmt.Printf("Failed reading memory in local player. %s", err)
			}
			var flags gorwmem.Data
			flags, err = dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MFFlags)), 0, gorwmem.UINT)
			if err != nil {
				fmt.Printf("Failed reading memory in flags. %s", err)
			}
			if (flags.Value.(uint32) & csgoFlOnGround) > 0 {
				if err = dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceJump)),
					gorwmem.Data{Value: csgoForceJump, DataType: gorwmem.INT}); err != nil {
					fmt.Printf("Failed reading memory in flags. %s", err)
				}
			}
		}
	}
}
