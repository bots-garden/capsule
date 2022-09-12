#!/bin/bash

# -----------------------------------------------
#  Tools
# -----------------------------------------------
function check_file() {
  file=$1
  if [ -f "${file}" ]; then
    echo "😃 ${file} exists."
  else
    echo "😡 ${file} does not exist."
		exit 1
  fi
}

eval $(cat vm.nats.config)

# check if certificates exist
#echo "🥇 checking certificates"
#check_file certs/${vm_domain}.crt
#check_file certs/${vm_domain}.key

# -----------------------------------------------
#  Create the primary node VM
# -----------------------------------------------
echo "🖥️ creating ${vm_name}"

multipass launch --name ${vm_name} \
--cpus ${vm_cpus} \
--mem ${vm_mem} \
--disk ${vm_disk} \
--cloud-init ./capsule.cloud-init.yaml

multipass mount certs ${vm_name}:certs

ip=$(multipass info ${vm_name} | grep IPv4 | awk '{print $2}')

multipass info ${vm_name}

multipass exec ${vm_name} -- sudo -- sh -c "echo \"${ip} ${vm_domain}\" >> /etc/hosts"

# install tools
multipass --verbose exec ${vm_name} -- bash <<EOF
sudo apt install unzip
EOF



# 🖐 add this to `hosts` file(s)
echo "${ip} ${vm_domain}" > config/caps.hosts.config

# 🖐 use this file to exchange data between VM creation script
# use: eval $(cat ../registry/workspace/export.registry.config)
target="config/export.caps.config"
echo "vm_name=\"${vm_name}\";" >> ${target}
echo "vm_domain=\"${vm_domain}\";" >> ${target}
echo "vm_ip=\"${ip}\";" >> ${target}

# -----------------------------------------------
#  Activate SSL
# -----------------------------------------------

multipass --verbose exec ${vm_name} -- sudo -- bash <<EOF
export LC_ALL="en_US.UTF-8"
export LC_CTYPE="en_US.UTF-8"

chmod 755 certs
EOF


echo "+-----------------------------------------------+"
echo "🖐️ update your /etc/hosts file with:"
cat config/caps.hosts.config
echo "+-----------------------------------------------+"
