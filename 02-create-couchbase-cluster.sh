PATH=$PATH:/opt/couchbase/bin

# ------------------------------------
# Create Couchbase cluster
# ------------------------------------
couchbase-cli cluster-init -c 127.0.0.1 \
--cluster-username admin \
--cluster-password ilovepandas \
--services data,index,query \
--cluster-ramsize 4096
