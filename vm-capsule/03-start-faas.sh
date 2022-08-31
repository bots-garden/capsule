#!/bin/bash
eval $(cat vm.capsule.config)
multipass --verbose exec ${vm_name} -- bash <<EOF
cd faas
./start.sh
EOF

