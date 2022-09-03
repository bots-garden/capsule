module github.com/bots-garden/capsule-faas-demo/hello

go 1.18

replace github.com/bots-garden/capsule/capsulemodule => ../../capsulemodule

replace github.com/bots-garden/capsule/commons => ../../commons

require github.com/bots-garden/capsule/capsulemodule v0.0.0-20220815065503-73f5c5e0cecb

require github.com/bots-garden/capsule/commons v0.0.0-20220903105536-e833f3d14593 // indirect
