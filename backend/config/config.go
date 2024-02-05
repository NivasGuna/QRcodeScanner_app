package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)


func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "YOUR_REDIS_ADDRESS", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
	} else {
		fmt.Printf("Connected to Redis: %v\n", pong)
	}

	return client
}

func GetNeo4jDriver() (neo4j.Driver, error) {
	uri := "YOUR_URI" // Replace with your Neo4j URI
	username := "YOUR_USERNAME"            // Replace with your Neo4j username
	password := "YOUR_PASSWORD"         // Replace with your Neo4j password

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Print("neo4j error")

	}

	return driver, nil
}
