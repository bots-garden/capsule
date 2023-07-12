# 🚀 Getting Started

## Use the Capsule HTTP server with Ngrok

With the `ngrok-go` library, you can use the Capsule HTTP server with Ngrok (you need to create an account on [Ngrok](https://ngrok.com/)). 


```bash
NGROK_AUTHTOKEN="${YOUR_NGROK_AUTHTOKEN}" \
./capsule-http \
--wasm=./functions/hello-world/hello-world.wasm \
--httpPort=6666
```

The ouput will be like this:

```bash
2023/05/18 11:25:36 💊 Capsule v0.4.1 🫑 [pepper] http server is listening on: 6666 🌍
2023/05/18 11:25:37 👋 Ngrok tunnel created: https://d298-88-173-112-231.ngrok-free.app
2023/05/18 11:25:37 🤚 Ngrok URL: /home/ubuntu/workspaces/capsule/capsule-http/ngrok.url
```

And you can access the wasm service with this url: `https://d298-88-173-112-231.ngrok-free.app` (the ngrok url is generated and different each time).


👋 You can retrieve the ngrok url in this file `ngrok.url`

If you own a Ngrok subscription, you can set your ngrok domain like this:

```bash
NGROK_DOMAIN="${YOUR_NGROK_DOMAIN}" \ # something like that "capsule.ngrok.dev"
NGROK_AUTHTOKEN="${YOUR_NGROK_AUTHTOKEN}" \
./capsule-http \
--wasm=./functions/hello-world/hello-world.wasm \
--httpPort=6666
```


!!! info "Ngrok and ngrok-go"
    - [Ngrok](https://ngrok.com/)
    - [ngrok-go](https://ngrok.com/blog-post/ngrok-go)
