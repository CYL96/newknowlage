package test

import (
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

func MyRedis(){
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	defer client.Close()
	pong,err := client.Ping().Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("Result:",pong)

	str := client.Set("str1","hello myboy",10*time.Second)
	fmt.Println(str)
	str1 := client.Get("str1")
	fmt.Println(str1.String())
	time.Sleep(11*time.Second)
	str1 = client.Get("str1")
	str2 := client.Exists("str1")
	fmt.Println(str2.String(),str1.String())
}

