module github.com/bots-garden/capsule/wasm_modules/reverse-proxy-demo/capsule-hola-yello

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-20220815100856-6b0da1ba4aad

require github.com/bots-garden/capsule/commons v0.0.0-20220903105536-e833f3d14593 // indirect
