package modules

// Once again credit to qhdwight for his basis

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/../../build -lwpiHal -lwpiutil -lstdc++ -ldl -lm -lFRC_NetworkCommunication -lembcanshim -lfpgalvshim -lRoboRIO_FRC_ChipObject -lvisa
// #include "hal.h"
import "C"
import (
	"os"
	"unsafe"
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

func Loader(L *lua.LState) int {
	
}