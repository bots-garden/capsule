const mqtt = require('mqtt')
const client  = mqtt.connect('mqtt://localhost:1883')

let config = {
  topic: "topic/sensor0",
  payload: "hello",
  clientId: "12345"
}

client.on('connect', _ => {
  console.log("connected")
  client.publish(config.topic, config.payload)
  client.end()
})
