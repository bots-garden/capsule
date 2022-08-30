# Multipass VM for tests
>
## Update the `vm.capsule.config` file

```
vm_name="capsdev-dev-devsecops-fun";
vm_domain="capsdev.devsecops.fun";
vm_cpus=3;
vm_mem="1G";
vm_disk="3GB";
```

If `vm_domain` equals `"capsdev.devsecops.fun"` and you want to enable HTTPS, you have to provide two files:
- `capsdev.devsecops.fun.crt`
- `capsdev.devsecops.fun.key`
Then you need to copy these two files to the `/certs` directory

## Create the VM

```bash
./create-vm.sh
```

When the VM is created, a `caps.hosts.config` is created in the `/config`. Copy its content to your `/etc/hosts` file.

## Connect to the VM with SSH

```bash
./shell-vm.sh
```

## Stop the VM

```bash
./stop-vm.sh
```

## Start the VM

```bash
./start-vm.sh
```

## Delete the VM

```bash
./destroy-vm.sh
```

## Capsule testing with the VM

### Build Capsule for the appropriate target

> For example, on a MacBook M1, the Multipass Ubuntu image architecture is arm:

```bash
cd ../capsule-launcher
env GOOS=linux GOARCH=arm64 go build -o capsule_arm64
cp capsule_arm64 ../vm-capsule/apps/capsule
rm capsule_arm64
```

### Install Capsule in the VM

Connect with SSH:
```bash
./shell-vm.sh
```

Then:
```bash
cd apps
sudo cp capsule /usr/local/bin
# and type the below command to check the install:
caps version
# you should get something like this: v0.1.7 🦑[squid]
```

### Launch Capsule with the HTTP mode

```bash
cd $HOME
MESSAGE="🖐 good morning 😄" \
capsule \
   -wasm=./src/functions/yo/yo.wasm \
   -mode=http \
   -crt=./certs/capsdev.devsecops.fun.crt \
   -key=./certs/capsdev.devsecops.fun.key \
   -httpPort=7070
```
> You should get something like: `💊 Capsule ( v0.1.7 🦑[squid] ) http server is listening on: 7070 🔐🌍`

Open https://capsdev.devsecops.fun:7070 with your favorite brother

