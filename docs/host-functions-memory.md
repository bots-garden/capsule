# Host functions

## Use memory cache

*`MemorySet`*
```golang
_, err := hf.MemorySet("message", "🚀 hello is started")
```

*`MemoryGet`*
```golang
value, err := hf.MemoryGet("message")
```

*`MemoryKeys`*
```golang
keys, err := hf.MemoryKeys()
// it will return an array of strings
if err != nil {
  hf.Log(err.Error())
}

for _, value := range keys {
  hf.Log("key: "+value)
}
```
