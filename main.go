package main

import (
	"log"
	"os"
	"robolua/modules"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	path_to_source := os.Args[1];

	L := lua.NewState();

	defer L.Close();

	// TODO: Load libraries

	modules.LoadModules(L);

	if err := L.DoFile(path_to_source); err != nil {
		log.Fatal("Error loading source file", "err", err);
	}

	robotMain := L.GetGlobal("main");

	if robotMain.Type() != lua.LTFunction {
		log.Fatal("main function not found");
	}


	modules.InitalizeHal();

	if err := L.CallByParam(lua.P{
		Fn: robotMain,
		NRet: 0,
	}); err != nil {
		log.Fatal("Error running main function", "err", err);
	}


}