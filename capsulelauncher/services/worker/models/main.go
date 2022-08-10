package models

import "reflect"

type RunningWasmModule struct {
	Pid       int
	Status    string // for future feature(s)
	LocalUrl  string
	RemoteUrl string
}

type Revision struct {
	Name string // optional

	WasmRegistryUrl string
	WasmModules     map[int]RunningWasmModule
}

type Function struct {
	Name      string
	Revisions map[string]Revision
}

func IsFunctionExist(functionName string, functions map[string]Function) bool {
	return !reflect.DeepEqual(functions[functionName], Function{})
}

func IsRevisionExist(functionName, revisionName string, functions map[string]Function) bool {
	return !reflect.DeepEqual(functions[functionName].Revisions[revisionName], Revision{})
}
