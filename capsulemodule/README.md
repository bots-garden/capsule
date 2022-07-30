# Capsule Module (wasm)

## Create exposed functions
> module function callable from the host

```golang

//export helloWorld
func helloWorld() (strPtrPosSize uint64) {
	strPtrPos, size := memory.GetStringPtrPositionAndSize("👋 hello world 🌍")
	return memory.PackPtrPositionAndSize(strPtrPos, size)
}

//export sayHello
func sayHello(strPtrPos, size uint32) (strPtrPosSize uint64) {
	name := helpers.GetStringParam(strPtrPos, size)
	pos, length := helpers.GetStringPtrPositionAndSize("👋 hello " + name)

	return helpers.PackPtrPositionAndSize(pos, length)
}
```
