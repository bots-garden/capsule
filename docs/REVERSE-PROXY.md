# Capsule Reverse Proxy


## Use the Capsule Reverse Proxy

You can use the **Capsule Reverse Proxy**. Then, you can call a function by its name:
```bash
http://localhost:8888/functions/hola
```
> *The reverse proxy will serve the **default** version of the `hola` function*

Or, you can use a revision of the function (for example, if you use several version of the function):
```bash
http://localhost:8888/functions/hola/orange
```
> - *The reverse proxy will serve the `orange` revision of the `hola` function*
> - *The `default` revision is the `default` version of the function (http://localhost:8888/functions/hola)*

To run the **Capsule Reverse Proxy**, run the below command:

```bash
./capsule-reverse-proxy \
   -config=./config.yaml \
   -backend="memory" \
   -httpPort=8888
```

With the Capsule Reverse Proxy, you gain an **API** that allows to define routes dynamically (in memory).

#### Registration API

##### Register a function (add a new route to a function)

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen", "revision": "default", "url": "http://localhost:5050"}'
```
> - This will add a new entry to the routes list with a `default` revision, with one url `http://localhost:5050`
> - You can call the function with this url: http://localhost:8888/function/morgen

The routes list (it's a map) will look like that:

```json
{
    "morgen": {
        "default": [
            "http://localhost:5050"
        ]
    }
}
```

You can create a new function registration with a named revision:

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen", "revision": "magenta", "url": "http://localhost:5051"}'
```
> - This will add a new entry to the routes list with a `magenta` revision, with one url `http://localhost:5051`
> - You can call the function with this url: http://localhost:8888/function/morgen/magenta


##### Remove the registration

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/registration \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"function": "morgen"}'
```


#### Revision API

##### Add a revision to the function registration

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/morgen/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "blue", "url": "http://localhost:5051"}'
```
> - The function already exists
> - The name of the function is set in the url `http://localhost:8888/memory/functions/:function_name/revision`

The routes list will look like that:

```json
{
    "morgen": {
        "blue": [
            "http://localhost:5051"
        ],
        "default": [
            "http://localhost:5050"
        ]
    }

}
```

##### Remove a revision from the function registration

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/morgen/revision \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"revision": "blue"}'
```

#### URL API

##### Add a URL to the revision of a function

```bash
curl -v -X POST \
  http://localhost:8888/memory/functions/morgen/blue/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:5053"}'
```
> - The revision already exists
> - The name of the function and of the revision are set in the url `http://localhost:8888/memory/functions/:function_name/:function_revision/url`

The routes list will look like that:

```json
{
    "morgen": {
        "blue": [
            "http://localhost:5051",
            "http://localhost:5053"
        ],
        "default": [
            "http://localhost:5050"
        ]
    }

}
```

> *A revision can be a set of URLs. In this case, the Capsule reverse-proxy will use randomly one of the URLs.*


##### Remove a URL from the function revision

```bash
curl -v -X DELETE \
  http://localhost:8888/memory/functions/morgen/blue/url \
  -H 'content-type: application/json; charset=utf-8' \
  -d '{"url": "http://localhost:5053"}'
```
