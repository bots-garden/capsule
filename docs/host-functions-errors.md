# Host functions

## Error Management
> 🖐🖐🖐 🚧 it's a work in progress (it's not implemented entirely)

*`GetExitError()` & `GetExitCode`*:
```golang
//export OnExit
func OnExit() {
	hf.Log("👋🤗 have a nice day")
	hf.Log("Exit Error: " + hf.GetExitError())
	hf.Log("Exit Code: " + hf.GetExitCode())
}
```
