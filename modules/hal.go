package modules

// Once again credit to qhdwight for his frc-go project, which this is based off of.

// #cgo CFLAGS: -I${SRCDIR}/includes
// #cgo LDFLAGS: -L${SRCDIR}/../../build -lwpiHal -lwpiutil -lstdc++ -ldl -lm -lFRC_NetworkCommunication -lembcanshim -lfpgalvshim -lRoboRIO_FRC_ChipObject -lvisa
// #include "hal.h"
import "C"
import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

const (
	FEnabled = 1 << iota
	FAutonomous
	FTest
	FEStop
	FFMSAttached
	FDSAttached
)

const (
	None = iota
	Disabled
	Autonomous
	Teleop
	Test
)

var exports = map[string]lua.LGFunction{
	"getTeamNumber": getTeamNumber,
}

func InitalizeHal() {
	ret := C.HAL_Initialize(500, 0)

	if ret == 0 { 
		log.Fatal("Failed to initialize HAL")
	}
}