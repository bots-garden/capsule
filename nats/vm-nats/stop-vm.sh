#!/bin/bash
eval $(cat vm.nats.config)
multipass stop ${vm_name}

