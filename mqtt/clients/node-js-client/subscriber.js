const mqtt = require('mqtt')
const client  = mqtt.connect('mqtt://localhost:1883')

let config = {
  topic: "topic/sensor0",
  clientId: "ABCDE"
}

client.on('connect', _ => {
  client.subscribe(config.topic, (error) => {
    console.log(error ? `😡 ${error.message}` : `😃 subscribed to ${config.topic} topic`)
  })
})

client.on('message',  (topic, message) => {
  console.log(message.toString())
})