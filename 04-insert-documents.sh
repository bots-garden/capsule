PATH=$PATH:/opt/couchbase/bin

# ------------------------------------
# Create documents
# ------------------------------------

script1='CREATE SCOPE `wasm-data`.data'
script2='CREATE COLLECTION `wasm-data`.data.docs'
script3='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key1", { "type" : "info", "name" : "this is an info" });'
script4='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key2", { "type" : "info", "name" : "this is another info" });'
script5='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key3", { "type" : "message", "name" : "👋 hello world 🌍" });'
script6='INSERT INTO `wasm-data`.data.docs (KEY, VALUE) VALUES ("key4", { "type" : "message", "name" : "👋 greetings 🎉" });'

cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script1"
cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script2"
cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script3"
cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script4"
cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script5"
cbq -u admin -p ilovepandas -e "http://localhost:8091" --script="$script6"

