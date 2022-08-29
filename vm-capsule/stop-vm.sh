#!/bin/bash
eval $(cat vm.capsule.config)
multipass stop ${vm_name}

