package process

import (
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/maxwelbm/gorwmem"
)

func Wall(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, _ := dm.GetModuleFromName("client.dll")
		glowManager, _ := dm.Read(uint(clientAddress)+offsets.Signatures.DwGlowObjectManager, gorwmem.UINT)
		for i := 1; i < 64; i++ {
			entity, _ := dm.Read((uint)(clientAddress)+(offsets.Signatures.DwEntityList+(uint)(i*0x10)), gorwmem.UINT)
			if entity.Value.(uint32) > 0 {
				entityTeamId, _ := dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MITeamNum, gorwmem.UINT)
				entityGlow, _ := dm.Read(uint(entity.Value.(uint32))+offsets.Netvars.MIGlowIndex, gorwmem.UINT)
				red := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x8)
				green := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0xC)
				blue := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x10)
				alpha := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x14)
				enable := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32)*0x38 + 0x28)

				if entityTeamId.Value.(uint32) == 2 {
					dm.Write(red, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT})
					dm.Write(green, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT})
					dm.Write(blue, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT})
					dm.Write(alpha, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT})
					dm.Write(enable, gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT})
				}
				if entityTeamId.Value.(uint32) == 3 {
					dm.Write(red, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT})
					dm.Write(green, gorwmem.Data{Value: float32(disableVal), DataType: gorwmem.FLOAT})
					dm.Write(blue, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT})
					dm.Write(alpha, gorwmem.Data{Value: float32(enableVal), DataType: gorwmem.FLOAT})
					dm.Write(enable, gorwmem.Data{Value: enableVal, DataType: gorwmem.UINT})
				}
			}
		}
	}
}
