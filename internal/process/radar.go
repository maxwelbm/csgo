package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/maxwelbm/gorwmem"
)

func Radar(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, err := dm.GetModuleFromName("client.dll")
		if err != nil {
			fmt.Printf("Failed reading module client.dll. %s", err)
		}
		for i := 1; i < 32; i++ {
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
				if entityTeamId.Value.(uint32) == 2 {
					if err = dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)),
						gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
				if entityTeamId.Value.(uint32) == 3 {
					if err = dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)),
						gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
			}
		}
	}
}
