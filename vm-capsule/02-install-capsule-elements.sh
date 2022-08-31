#!/bin/bash

eval $(cat vm.capsule.config)

multipass --verbose exec ${vm_name} -- bash <<EOF
cd apps
sudo cp capsule /usr/local/bin
sudo cp capsule-registry /usr/local/bin
sudo cp capsule-reverse-proxy /usr/local/bin
sudo cp capsule-worker /usr/local/bin
EOF

#multipass --verbose exec ${vm_name} -- sudo -- bash <<EOF
#
#EOF
