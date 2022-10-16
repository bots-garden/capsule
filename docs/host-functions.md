#Host functions

**Capsule** offers some capabilities to the wasm modules by providing some "host functions":

## Print a message

```golang
hf.Log("👋 Hello World 🌍")
```

## Read and Write files

```golang
txt, err := hf.ReadFile("about.txt")
if err != nil {
    hf.Log(err.Error())
}
hf.Log(txt)

newFile, err := hf.WriteFile("hello.txt", "👋 HELLO WORLD 🌍")
if err != nil {
    hf.Log(err.Error())
}
hf.Log(newFile)
```

## Read value of the environment variables

```golang
message, err := hf.GetEnv("MESSAGE")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("MESSAGE=" + message)
}
```

## Make HTTP requests

*`GET`*
```golang
ret, err := hf.Http("https://httpbin.org/get", "GET", headers, "")
if err != nil {
    hf.Log("😡 error:" + err.Error())
} else {
    hf.Log("📝result: " + ret)
}
```

*`POST`*
```golang
headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

ret, err := hf.Http("https://httpbin.org/post", "POST", headers, "👋 hello world 🌍")
if err != nil {
    hf.Log("😡 error:" + err.Error())
} else {
    hf.Log("📝result: " + ret)
}
```

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

for key, value := range keys {
  hf.Log(key+":"+value)
}
```

## Make Redis queries
> 🚧 this is a work in progress

You need to run **Capsule** with these two environment variables:
```bash
REDIS_ADDR="localhost:6379"
REDIS_PWD=""
```

*`SET`*
```golang
// add a key, value
res, err := hf.RedisSet("greetings", "Hello World")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("Value: " + res)
}
```

*`GET`*
```golang
// read the value
res, err := hf.RedisGet("greetings")
if err != nil {
    hf.Log(err.Error())
} else {
    hf.Log("Value: " + res)
}
```

*`KEYS`*
```golang
legion, err := hf.RedisKeys("bob*")
if err != nil {
    hf.Log(err.Error())
}

for _, bob := range legion {
    hf.Log(bob)
}
```

## Make CouchBase N1QL Query

You need to run **Capsule** with these four environment variables:
```bash
COUCHBASE_CLUSTER="couchbase://localhost"
COUCHBASE_USER="admin"
COUCHBASE_PWD="ilovepandas"
COUCHBASE_BUCKET="wasm-data"
```

```golang
bucketName, _ := hf.GetEnv("COUCHBASE_BUCKET")
query := "SELECT * FROM `" + bucketName + "`.data.docs"

jsonStringArray, err := hf.CouchBaseQuery(query)
```

## Nats

*`NatsPublish(subject string, message string)`*: publish a message on a subject
```golang
_, err := hf.NatsPublish("notify", "it's a wasm module here")
```
> You must use the `"nats"` mode of **Capsule** as the NATS connection is defined at the start of **Capsule** and shared with the WASM module:

```bash
capsule \
 -wasm=../wasm_modules/capsule-nats-subscriber/hello.wasm \
 -mode=nats \
 -natssrv=nats.devsecops.fun:4222 \
 -subject=ping
```

*`NatsReply(message string, timeout uint32)`*: publish a message on the current subject and wait for an answer
```golang
_, err := hf.NatsReply("it's a wasm module here", 10)
```
> You must use the `"nats"` mode of **Capsule** as the NATS connection and the subscription are defined at the start of **Capsule** and shared with the WASM module.



*`NatsGetSubject()`*: get the subject listened by the **Capsule** launcher
```golang
hf.Log("👂Listening on: " + hf.NatsGetSubject())
```


*`NatsGetServer()`*: get the connected NATS server
```golang
hf.Log("👋 NATS server: " + hf.NatsGetServer())
```

*`NatsConnectPublish(server string, subject string, message string)`*: connect to a NATS server and send a message on a subject
```golang
_, err := hf.NatsConnectPublish("nats.devsecops.fun:4222", "ping", "🖐 Hello from WASM with Nats 💜")
```
> You can use this function with all the running modes of **Capsule**

*`NatsConnectPublish(server string, subject string, message string, timeout uint32)`*: connect to a NATS server and send a message on a subject
```golang
answer, err := hf.NatsConnectRequest("nats.devsecops.fun:4222", "notify", "👋 Hello World 🌍", 1)
```
> You can use this function with all the running modes of **Capsule**


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
