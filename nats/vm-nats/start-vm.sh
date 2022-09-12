#!/bin/bash
eval $(cat vm.nats.config)
multipass start ${vm_name}
