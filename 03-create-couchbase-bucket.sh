PATH=$PATH:/opt/couchbase/bin

# ------------------------------------
# Create Couchbase bucket
# ------------------------------------
couchbase-cli bucket-create -c 127.0.0.1:8091 \
--username admin \
--password ilovepandas \
--bucket wasm-data \
--bucket-type couchbase \
--bucket-ramsize 1024


