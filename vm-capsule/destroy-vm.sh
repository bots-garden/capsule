#!/bin/bash
eval $(cat vm.capsule.config)
multipass delete ${vm_name}
multipass purge

rm  config/caps.hosts.config
rm  config/export.caps.config

