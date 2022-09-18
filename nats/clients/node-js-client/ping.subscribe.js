// subscribe to "ping" message

import { connect, StringCodec } from "nats";

console.log("Connecting to the server...")
//const nc = await connect({ servers: ["192.168.64.18:4222"] });
const nc = await connect({ servers: ["nats.devsecops.fun:4222"] });

// create a codec
const sc = StringCodec();
// create a simple subscriber and iterate over messages
// matching the subscription
const sub = nc.subscribe("ping");
(async () => {
    for await (const m of sub) {
        console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
    }
    console.log("subscription closed");
})();
