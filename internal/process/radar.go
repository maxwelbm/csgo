package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/maxwelbm/gorwmem"
)

const (
	enableVal  = 1
	disableVal = 0
)

func Radar(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, _ := dm.GetModuleFromName("client.dll")
		for i := 1; i < 32; i++ {
			entity, _ := dm.Read((uint)(clientAddress)+(offsets.Signatures.DwEntityList+(uint)(i*0x10)), gorwmem.UINT)
			if entity.Value.(uint32) > 0 {
				entityTeamId, _ := dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MITeamNum, gorwmem.UINT)
				if entityTeamId.Value.(uint32) == 2 {
					if err := dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)),
						gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
				if entityTeamId.Value.(uint32) == 3 {
					if err := dm.Write(uint(entity.Value.(uint32)+uint32(offsets.Netvars.MBSpotted)),
						gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT}); err != nil {
						fmt.Printf("Failed writing memory. %s", err)
						continue
					}
				}
			}
		}
	}
}
