import { WebSocketServer } from 'ws';
import { DetectionDto } from './model/Detection';
const redis = require("redis");
const droneChannel = "droneMove" // hardcoded but we can put this in a yaml file
const subscriber = redis.createClient({
  port:6379,
  host:"localhost"
});
subscriber.subscribe(droneChannel)
const wss = new WebSocketServer({ port: 8080 });


// // Sydney International Airport example
let droneCoords: DetectionDto = {
  Latitude: -33.937687,
  Longitude: 151.19189864
};

// update drone coordinates
const updateDrone = (droneCoords): DetectionDto => {
  return {
    Latitude: droneCoords.Latitude,
    Longitude: droneCoords.Longitude
  };
};

// listen for new incoming coordinates from producer
subscriber.on("message",(channel,message) => {
  console.log("Subscriber received message in channel '" + channel + "': " + message);
  let payload = JSON.parse(message)
  droneCoords = updateDrone(payload)
  
})


wss.on('connection', (ws) => {
  setInterval(() => {
    ws.send(JSON.stringify(droneCoords));
  }, 2000);
});
