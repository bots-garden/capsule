import { connect, StringCodec } from "nats";

console.log("Connecting to the server...")
const nc = await connect({ servers: ["nats.devsecops.fun:4222"] });

// create a codec
const sc = StringCodec();

nc.publish("ping", sc.encode("Hello 👋"));
nc.publish("ping", sc.encode("Morgen 😃"));

await nc.drain();
