#!/bin/bash
eval $(cat vm.capsule.config)
multipass shell ${vm_name}

