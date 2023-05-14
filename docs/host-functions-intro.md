# Host functions

> 🚧 this is a work in progress, each host function will be detailed with samples.

The **Capsule** applications (and the Capsule HDK) offer some capabilities to the wasm modules by providing some "host functions".

This is the list of the available host functions:

- Print a message: `Print(message string)`, usage: `capsule.Print("👋 Hello Worls 🌍")`
- Log a message: `Log(message string)`, usage: `capsule.Log("😡 something wrong")`
- Get the value of an environment variable: `GetEnv(variableName string) string`, usage: `capsule.GetEnv("MESSAGE")`
- Read a text file: `ReadFile(filePath string) ([]byte, error)`, usage: `data, err := capsule.ReadFile("./hello.txt")`
- Write a content to a text file: `WriteFile(filePath string, content []byte) error`, usage: `err := capsule.WriteFile("./hello.txt", []byte("👋 Hello World! 🌍"))`
- Make an HTTP request: `HTTP(request HTTPRequest) (HTTPResponse, error)`, usage: `respJSON, err := capsule.HTTP(capsule.HTTPRequest{})`, see the ["hey-people" sample](https://github.com/bots-garden/capsule/blob/main/capsule-cli/functions/hey-people/main.go#L15)
- Memory Cache: see the ["mem-db" sample](https://github.com/bots-garden/capsule/blob/main/capsule-cli/functions/mem-db/main.go)
  - `CacheSet(key string, value []byte) []byte`
  - `CacheGet(key string) ([]byte, error)`
  - `CacheDel(key string) []byte`
  - `CacheKeys(filter string) ([]string, error)` (right now, you can only use this filter: `*`)
- Redis Cache: see the ["redis-db" sample](https://github.com/bots-garden/capsule/blob/main/capsule-cli/functions/redis-db/main.go)
  - `RedisSet(key string, value []byte) ([]byte, error)`
  - `RedisGet(key string) ([]byte, error)`
  - `RedisDel(key string) ([]byte, error)`
  - `RedisKeys(filter string) ([]string, error)`
