package modules

// #cgo CFLAGS: -I${SRCDIR}/includes
// #cgo LDFLAGS: -L${SRCDIR}/../../build -lwpiHal -lwpiutil -lstdc++ -ldl -lm -lFRC_NetworkCommunication -lembcanshim -lfpgalvshim -lRoboRIO_FRC_ChipObject -lvisa
// #include "hal.h"
import "C"
import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

var globals = map[string]lua.LGFunction{
	"print": print,
}

func LoadGlobals(L *lua.LState) int {
	for name, fn := range globals { 
		L.SetGlobal(name, L.NewFunction(fn))
	}

	return 1
}

func print(L *lua.LState) int {
	n := L.GetTop()
	for i := 1; i <= n; i++ {
		L.Push(L.Get(i))
		log.Print(L.ToString(-1))

		C.HAL_SendConsoleLine(L.ToString(-1))
		L.Pop(1)
	}
	return 0
}

func getTeamNumber(L *lua.LState) int {
	L.Push(lua.LNumber(C.HAL_GetTeamNumber()))
	return 1
}