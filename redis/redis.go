package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"my-auth-api/models"
)

var ctx = context.Background()
var rdb *redis.Client

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	fmt.Println("Connected to Redis")
}

func SaveUser(user models.User) error {
	return rdb.HSet(ctx, user.Username, "password", user.Password).Err()
}

func ValidateUser(user models.User) bool {
	storedPassword, err := rdb.HGet(ctx, user.Username, "password").Result()
	if err != nil || storedPassword != user.Password {
		return false
	}
	return true
}
