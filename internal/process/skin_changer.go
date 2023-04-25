package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/MaxwelMazur/csboost/internal/model/weapon"
	"github.com/maxwelbm/gorwmem"
	"time"
)

func SkinChanger(dm *gorwmem.DataManager, offsets *model.OffSet) {
	client, err := dm.GetModuleFromName("client.dll")
	if err != nil {
		fmt.Printf("Failed reading module client.dll. %s", err)
	}

	engine, err := dm.GetModuleFromName("engine.dll")
	if err != nil {
		fmt.Printf("Failed reading module engine.dll. %s", err)
	}

	for {
		localPlayer, _ := dm.Read(uint(client)+offsets.Signatures.DwLocalPlayer, 0, gorwmem.UINT)

		weapons, _ := dm.Read(uint(localPlayer.Value.(uint32))+offsets.Netvars.MHMyWeapons, 8, gorwmem.ARRAY)
		listWeapons := weapons.Value.([]byte)
		for _, v := range listWeapons {
			time.Sleep(2 * time.Millisecond)

			weapon, _ := dm.Read(uint(client)+offsets.Signatures.DwEntityList+uint((uintptr(v)&0xFFF)*0x10-0x10), 0, gorwmem.UINT)
			if weapon.Value == 0 {
				continue
			}

			itemCurrent, _ := dm.Read(uint(weapon.Value.(uint32))+offsets.Netvars.MIItemDefinitionIndex, 0, gorwmem.UINT)
			paint := GetWeaponPaint(itemCurrent.Value.(uint32))

			if paint != 0 {
				shouldUpdate, _ := dm.Read(uint(weapon.Value.(uint32))+offsets.Netvars.MNFallbackPaintKit, 0, gorwmem.UINT)
				shouldUpdateGo := false
				if shouldUpdate.Value != paint {
					shouldUpdateGo = true
				}

				dm.Write(uint(weapon.Value.(uint32))+offsets.Netvars.MIItemIDHigh, gorwmem.Data{Value: int32(-1), DataType: gorwmem.INT})
				dm.Write(uint(weapon.Value.(uint32))+offsets.Netvars.MNFallbackPaintKit, gorwmem.Data{Value: int32(paint), DataType: gorwmem.INT})
				dm.Write(uint(weapon.Value.(uint32))+offsets.Netvars.MFlFallbackWear, gorwmem.Data{Value: float32(0.1), DataType: gorwmem.FLOAT})

				if !shouldUpdateGo {
					stateClient, _ := dm.Read(uint(engine)+offsets.Signatures.DwClientState, 0, gorwmem.UINT)
					dm.Write(uint(stateClient.Value.(uint32))+0x174, gorwmem.Data{Value: int32(-1), DataType: gorwmem.INT})
				}
			}
		}
	}
}

func GetWeaponPaint(itemDefinition uint32) int {
	weaponPaints := map[uint32]int{
		weapon.DEAGLE:       711,
		weapon.GLOCK:        38,
		weapon.AK47:         490,
		weapon.AWP:          344,
		weapon.USP_SILENCER: 653,
	}
	return weaponPaints[itemDefinition]
}
