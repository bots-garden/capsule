#!/bin/bash
eval $(cat vm.capsule.config)
multipass start ${vm_name}
