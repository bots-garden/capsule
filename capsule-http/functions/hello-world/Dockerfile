FROM botsgarden/capsule-http-linux-arm64:0.3.9
COPY hello-world.wasm .
EXPOSE 8080
CMD ["/capsule-http", "--wasm=./hello-world.wasm", "--httpPort=8080"]

