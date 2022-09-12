# Multipass VM for tests
> - THIS IS A WIP 🚧
> - Update the `vm.nats.config` file

```bash
sudo apt update
sudo apt install snapd
sudo systemctl status snapd.socket
```
https://github.com/nats-io/nats-server/releases/download/v2.9.0/nats-server-v2.9.0-arm64.deb

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

