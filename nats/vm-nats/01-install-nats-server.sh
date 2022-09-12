#!/bin/bash

eval $(cat vm.nats.config)

multipass --verbose exec ${vm_name} -- bash <<EOF
version="2.9.0"
os="linux-arm64"
curl -L https://github.com/nats-io/nats-server/releases/download/v${version}/nats-server-v${version}-${os}.zip -o nats-server.zip

#curl -L  https://github.com/nats-io/nats-server/releases/download/v2.9.0/nats-server-v2.9.0-linux-arm64.zip -o nats-server.zip
unzip nats-server.zip -d nats-server
sudo cp nats-server/nats-server-v${version}-${os}/nats-server /usr/bin
# sudo cp nats-server/nats-server-v2.9.0-linux-arm64/nats-server /usr/bin
rm nats-server.zip
rm -rf nats-server
EOF

#multipass --verbose exec ${vm_name} -- sudo -- bash <<EOF
#
#EOF
