#!/bin/bash
eval $(cat vm.nats.config)
multipass --verbose exec ${vm_name} -- bash <<EOF
nats-server --signal stop
EOF

