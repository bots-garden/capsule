# Worker

> We can have several worker

## Start a Worker

```bash
capsule-worker \
   -reverseProxy=http://localhost:8888 \
   -backend=memory \
   -capsulePath=capsule \
   -httpPortCounter=10000 \
   -httpPort=9999
```

### Pre-requisites

> **Start the registry**:
```bash
capsule-registry \
   -files="/functions" \
   -httpPort=4999
```

> **Start the reverse-proxy**:
```bash
capsule-reverse-proxy \
   -backend="memory" \
   -httpPort=8888
```


## Worker API

### Deployment

> **Start a function**
```bash
curl -v -X POST \
http://localhost:9999/functions/deploy \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current",
    "downloadUrl": "http://localhost:4999/k33g/hola/0.0.0/hola.wasm",
    "envVariables": {
        "MESSAGE": "🟣 Current revision of Hola",
        "TOKEN": "DFRE46H5K"
    }
}
EOF
```
> - You can start the same function with the same revision: it's like scaling

> Output:
```bash
{
  "code":"FUNCTION_DEPLOYED","
  function":"hola",
  "localUrl":"http://localhost:10001",
  "message":"Function deployed",
  "remoteUrl":"http://localhost:8888/functions/hola/current",
  "revision":"current"
}
```
> - You can access the function with the remote url (`functions/<function_name>/<revision_name>`)

### Default Revision

You can have several revisions for one function, but you can decide that one of the revisions is the default one (that means you can access the function with this URL: `functions/<function_name>`, without the revision name)

> **Set the default revision**
```bash
curl -v -X POST \
http://localhost:9999/functions/set_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current"
}
EOF
```
> - The `current` revision of `hola` becomes the default revision

> Output:
```bash
{
  "code":"DEFAULT_REVISION_REGISTERED",
  "revision":"current",
  "function":"hola",
  "message":"Default revision registered",
  "url":"http://localhost:8888/functions/hola"
}
```
> - You can reach the function with:
>   - http://localhost:8888/functions/hola
>   - http://localhost:8888/functions/hola/current
>   - http://localhost:8888/functions/hola/default


> **Remove the default revision (without undeploy the function)** *(then you can set again the default revision to another revision of the function)*
```bash
curl -v -X DELETE \
http://localhost:9999/functions/remove_default_revision \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola"
}
EOF
```

> Output:
```bash
{
  "code":"DEFAULT_REVISION_REMOVED",
  "function":"hola",
  "message":"Default revision removed"
}
```

> **Undeploy a revision** *(stop the function process and remove the revision)*
```bash
curl -v -X DELETE \
http://localhost:9999/functions/revisions/deployments \
-H "Expect:" \
-H 'content-type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
    "function": "hola",
    "revision": "current"
}
EOF
echo ""
```

> Output:
```bash
{
  "code":"REVISION_DEPLOYMENT_REMOVED",
  "function":"hola",
  "message":"revision deployment removed (all processes killed)",
  "revision":"current"
}
```

## Get some information from the worker and the reverse-proxy

> **Get the list of the started functions on the worker**:
```bash
curl http://localhost:9999/functions/list
```

> Output:
```bash
{
    "hola": {
        "blue": {
            "isDefaultRevision": "🚧 to be implemented",
            "wasmModules": {
                "10028": {
                    "envVariables": {
                        "MESSAGE": "🔵 Blue revision of Hola",
                        "TOKEN": "this is not a header token"
                    },
                    "localUrl": "http://localhost:10003",
                    "remoteUrl": "http://localhost:8888/functions/hola/blue"
                }
            },
            "wasmRegistryUrl": "http://localhost:4999/k33g/hola/0.0.1/hola.wasm"
        },
        "current": {
            "isDefaultRevision": "🚧 to be implemented",
            "wasmModules": {
                "10024": {
                    "envVariables": {
                        "MESSAGE": "1️⃣🟣 Current revision of Hola",
                        "TOKEN": "this is not a header token"
                    },
                    "localUrl": "http://localhost:10001",
                    "remoteUrl": "http://localhost:8888/functions/hola/current"
                },
                "10026": {
                    "envVariables": {
                        "MESSAGE": "2️⃣🟣 Current revision of Hola",
                        "TOKEN": "this is not a header token"
                    },
                    "localUrl": "http://localhost:10002",
                    "remoteUrl": "http://localhost:8888/functions/hola/current"
                }
            },
            "wasmRegistryUrl": "http://localhost:4999/k33g/hola/0.0.0/hola.wasm"
        },
        "green": {
            "isDefaultRevision": "🚧 to be implemented",
            "wasmModules": {
                "10030": {
                    "envVariables": {
                        "MESSAGE": "🟢 Green revision of Hola",
                        "TOKEN": "this is not a header token"
                    },
                    "localUrl": "http://localhost:10004",
                    "remoteUrl": "http://localhost:8888/functions/hola/green"
                }
            },
            "wasmRegistryUrl": "http://localhost:4999/k33g/hola/0.0.2/hola.wasm"
        }
    }
}
```

> **Get the list of the revisions registered to the reverse-proxy**:
```bash
curl http://localhost:8888/memory/functions/list
```
> `memory`: if you use the `memory` backend (right now, it's the only one)

> Output:
```bash
{
    "hola": {
        "blue": [
            "http://localhost:10003"
        ],
        "current": [
            "http://localhost:10001",
            "http://localhost:10002"
        ],
        "default": [
            "http://localhost:10001",
            "http://localhost:10002"
        ],
        "green": [
            "http://localhost:10004"
        ]
    }
}
```
