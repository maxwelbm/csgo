package process

import (
	"fmt"
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/jamesmoriarty/gomem"
	"github.com/maxwelbm/gorwmem"
	"math/rand"
	"time"
)

const (
	VkShift         = 0x10 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	VkControl       = 0x11
	CsgoForceAttack = 0x6
)

func Trigger(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		rtNum := reactionTime()
		fmt.Printf("\033[2K\rTempo gerado aleatoriamente: %d milisegundos", rtNum)
		time.Sleep(time.Duration(rtNum) * time.Millisecond)
		clientAddress, _ := dm.GetModuleFromName("client.dll")
		if gomem.IsKeyDown(VkShift) || gomem.IsKeyDown(VkControl) {
			localPlayer, _ := dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), gorwmem.UINT)
			entity, _ := dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MICrosshairId)), gorwmem.UINT)
			if entity.Value.(uint32) > 0 && entity.Value.(uint32) <= 64 {
				dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceAttack)), gorwmem.Data{Value: CsgoForceAttack, DataType: gorwmem.INT})
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
