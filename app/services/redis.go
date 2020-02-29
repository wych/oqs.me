package services

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"oqs.me/config"
)

var redisClient *redis.Client

func RedisInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("redis client init error")
	}
}

func IssueID() (uint, error) {
	i, err := redisClient.Incr("oqs:maxid").Result()
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func SetOQSID(id uint) error {
	err := redisClient.Set("oqs:maxid", id, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func MaxOQSID() uint {
	i, _ := redisClient.Get("oqs:maxid").Uint64()
	return uint(i)
}

func makeRecordCache(record string, realURL string) error {
	err := redisClient.Set(record, realURL, time.Minute*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func readRecordCache(record string) (string, error) {
	r, err := redisClient.Get(record).Result()
	if err != nil {
		return "", err
	}
	redisClient.Expire(record, time.Minute*30)
	return r, nil
}
