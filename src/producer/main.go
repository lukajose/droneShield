package main

import (
	"context"
	"log"
	"producer/droneMovement"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {

	// set up redis
	rds := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // connect to container
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
	// ping and check ok
	_, err := rds.Ping().Result()
	// retry redis connection
	for err != nil {
		log.Println("Reconnecting in 5s to redis")
		time.Sleep(5 * time.Second)
		rds = redis.NewClient(&redis.Options{
			Addr:     "redis:6379", // connect to container
			Password: "",           // no password set
			DB:       0,            // use default DB
		})
		_, err = rds.Ping().Result()
	}

	// Sydney International Airport example as starting point
	drone := droneMovement.DroneCoords{
		Latitude:  -33.937687,
		Longitude: 151.19189864,
	}

	log.Println("Connected to redis")
	log.Println("Running drone coords producer")
	drone.Start(rds)

}
