## How to run 

I dockerized almost everything, just didnt add the server so if you have docker-compose simply do, if for any reason you dont want to use docker-compose you can just change 
the host in main.go to be localhost and run manually.

```
docker-compose up 

```

Then run the server to connect the subscriber to redis

```
npm start

```

Finally if you want to connect to the websocket service I wrote a quick client websocket simply run 

```
node client.js

```

I only did two trajectory functions, a simple circle and a zig zag but more can be added accordingly using different math functions.

