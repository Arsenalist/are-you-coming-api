package main

import "github.com/go-redis/redis"

func Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18025.c44.us-east-1-2.ec2.cloud.redislabs.com:18025",
		Password: "9dtBwQvPM9cuiACjHoiRSJw4M0krORPa",
	})
	return client
}

