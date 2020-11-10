package main

import "github.com/go-redis/redis"

func main()  {
	data := "adsfadsfdsafsdf"
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := RedisClient.Publish("message", data).Err()
	if err != nil {
		return
	}
}
