## How to run 

Redis and the producer have been dockerized for you, just the server was left out, so if you have docker-compose simply do, 

```
docker-compose up 
```
If for any reason you dont want to use docker-compose you can just change 
the host in main.go to be localhost and run manually.

Now, to run the microservice consumer do, 

```
npm start
```

Finally if you want to connect to the websocket service I wrote a quick client websocket under client directory simply run,

```
node client.js
```

I only did two trajectory functions, a simple circle and a zig zag but more can be added accordingly using different math functions.

