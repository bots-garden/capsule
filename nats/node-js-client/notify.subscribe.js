// subscribe to "notify" message

import { connect, StringCodec } from "nats";

const servers = [
    {},
    { servers: ["192.168.64.18:4442"] }, // or ["localhost:4442"]
]

// { servers: ["nats.devsecops.fun:4442"] }, // or ["localhost:4442"]
//

console.log("Connecting to the first server...")
// to create a connection to a nats-server:
//const nc = await connect({ servers: ["192.168.64.18:4222"] });
const nc = await connect({ servers: ["nats.devsecops.fun:4222"] });


// create a codec
const sc = StringCodec();
// create a simple subscriber and iterate over messages
// matching the subscription
const sub = nc.subscribe("notify");
(async () => {
    for await (const m of sub) {
        console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
    }
    console.log("subscription closed");
})();
