package droneMovement

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/go-redis/redis"
)

type DroneCoords struct {
	Latitude  float64
	Longitude float64
}

const channel = "droneMove" // we could read this from yaml file
const radius = 10           // could be less or more depending on how wide we want out circle
const angleInterval = .2    // angle can be less for more precision

/* Creates the angle sequence for the circle trajectory and zig zag*/
func angleSeq() func() float64 {
	angle := 0.0
	return func() float64 {
		angle += angleInterval
		if angle > 2*math.Pi {
			// start again
			angle = angleInterval
		}
		return angle
	}
}

// update coordinates of a drone to continuosly do circles
func (d *DroneCoords) MoveCircle(rds *redis.Client, angle float64) error {

	droneMove := DroneCoords{
		Latitude:  d.Latitude + radius*math.Cos(angle),
		Longitude: d.Longitude + radius*math.Sin(angle),
	}
	//jsonify
	b, err := json.Marshal(droneMove)
	payload := string(b)
	log.Println("sending:", payload)
	if err == nil {
		rds.Publish(channel, payload)
		return nil
	}
	return fmt.Errorf("something went wrong stringifying struct")
}

// update coordinates of drone to continuosly do zig zags
func (d *DroneCoords) MoveZigZag(rds *redis.Client) error {

	droneMove := DroneCoords{
		Latitude:  d.Latitude,           //move right
		Longitude: math.Sin(d.Latitude), // do a sin function
	}
	d.Latitude += 0.00009173 // from example move right given
	// jsonify
	b, err := json.Marshal(droneMove)
	payload := string(b)
	log.Println("sending:", payload)
	if err == nil {
		rds.Publish(channel, payload)
		return nil
	}
	return fmt.Errorf("something went wrong stringifying struct")

}

/** Our start function that will execute all flying plans**/
func (d *DroneCoords) Start(rds *redis.Client) {
	angleInt := angleSeq()
	for {
		// here we can add different moves
		angle := angleInt()
		d.MoveCircle(rds, angle)
		time.Sleep(2 * time.Second)
	}
}
