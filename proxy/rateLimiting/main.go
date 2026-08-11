package main

import (
	"fmt"
	apiservice "rateLimiting/apiService"
	"time"
)

func main() {
	api := apiservice.NewRateLimiter()
	fmt.Println(api.Request("/users"))
	fmt.Println(api.Request("/orders"))
	fmt.Println(api.Request("/products"))
	fmt.Println(api.Request("/inventory"))
	fmt.Println("Sleeping for 10 second...")
	time.Sleep(10 * time.Second)
	fmt.Println(api.Request("/users"))
	fmt.Println(api.Request("/orders"))
	fmt.Println(api.Request("/products"))
	fmt.Println(api.Request("/inventory"))
}
