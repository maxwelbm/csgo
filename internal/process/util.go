package process

const (
	vkSpace         = 0x20   // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	csgoFlOnGround  = 1 << 0 // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
	csgoForceJump   = 0x6    // https://github.com/ValveSoftware/source-sdk-2013/blob/0d8dceea4310fde5706b3ce1c70609d72a38efdf/sp/src/game/shared/sdk/sdk_playeranimstate.cpp#L517
	enableVal       = 1
	disableVal      = 0
	vkShift         = 0x10 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	vkControl       = 0x11
	csgoForceAttack = 0x6
)
