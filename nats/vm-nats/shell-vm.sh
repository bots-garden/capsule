#!/bin/bash
eval $(cat vm.nats.config)
multipass shell ${vm_name}

