package process

import (
	"github.com/MaxwelMazur/csboost/internal/model"
	"github.com/jamesmoriarty/gomem"
	"github.com/maxwelbm/gorwmem"
)

const (
	VkSpace        = 0x20   // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	CsgoFlOnGround = 1 << 0 // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
	CsgoForceJump  = 0x6    // https://github.com/ValveSoftware/source-sdk-2013/blob/0d8dceea4310fde5706b3ce1c70609d72a38efdf/sp/src/game/shared/sdk/sdk_playeranimstate.cpp#L517
)

func BHop(dm *gorwmem.DataManager, offsets *model.OffSet) {
	for {
		clientAddress, _ := dm.GetModuleFromName("client.dll")
		if gomem.IsKeyDown(VkSpace) {
			localPlayer, _ := dm.Read(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwLocalPlayer)), gorwmem.UINT)
			flags, _ := dm.Read(uint(localPlayer.Value.(uint32)+(uint32)(offsets.Netvars.MFFlags)), gorwmem.UINT)
			if (flags.Value.(uint32) & CsgoFlOnGround) > 0 {
				dm.Write(uint((uint32)(clientAddress)+(uint32)(offsets.Signatures.DwForceJump)), gorwmem.Data{Value: CsgoForceJump, DataType: gorwmem.INT})
			}
		}
	}
}
