package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/maxwelbm/gorwmem"
)

func Wall(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, err := dm.GetModuleFromName("client.dll")
		if err != nil {
			fmt.Printf("Failed reading module client.dll. %s", err)
		}
		var glowManager gorwmem.Data
		glowManager, err = dm.Read(uint(clientAddress)+offsets.Signatures.DwGlowObjectManager, gorwmem.UINT)
		if err != nil {
			fmt.Printf("Failed reading memory glowManager. %s", err)
		}
		for i := 1; i < 64; i++ {
			var entity gorwmem.Data
			entity, err = dm.Read((uint)(clientAddress)+(offsets.Signatures.DwEntityList+(uint)(i*0x10)), gorwmem.UINT)
			if err != nil {
				fmt.Printf("Failed reading memory entity. %s", err)
			}
			if entity.Value.(uint32) > 0 {

				var entityTeamId gorwmem.Data
				entityTeamId, err = dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MITeamNum, gorwmem.UINT)
				if err != nil {
					fmt.Printf("Failed reading memory entityTeamId. %s", err)
				}

				var entityGlow gorwmem.Data
				entityGlow, err = dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MIGlowIndex, gorwmem.UINT)
				if err != nil {
					fmt.Printf("Failed reading memory entityGlow. %s", err)
				}

				red := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x8)
				green := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0xC)
				blue := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x10)
				alpha := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x14)
				enable := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x28)

				if entityTeamId.Value.(uint32) == 2 {
					if err = dm.Write(red, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(green, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(blue, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(alpha, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(enable, gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
				if entityTeamId.Value.(uint32) == 3 {
					if err = dm.Write(red, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(green, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(blue, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(alpha, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
					if err = dm.Write(enable, gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
			}
		}
	}
}
