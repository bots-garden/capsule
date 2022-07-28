# ------------------------------------
# Install Couchbase
# ------------------------------------
curl -O https://packages.couchbase.com/releases/couchbase-release/couchbase-release-1.0-amd64.deb
sudo dpkg -i ./couchbase-release-1.0-amd64.deb
sudo apt-get update
sudo apt-get install couchbase-server-community -y
rm ./couchbase-release-1.0-amd64.deb
#couchbase-cli cluster-init -c 127.0.0.1 --cluster-username admin \
#--cluster-password admin --services data --cluster-ramsize 4096
