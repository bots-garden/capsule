#!/bin/bash

# bash -c "exec -a <MyProcessName> <Command>"

cd ./capsule-hello
./build.sh

cd ../capsule-hey
./build.sh

cd ../capsule-hola
./build.sh

cd ../capsule-hola-orange
./build.sh

cd ../capsule-hola-yellow
./build.sh

