package main

import (
	"fmt"
	"time"
)

func main() {
	var amount int
	fmt.Println("enter ye time ye cont")
	fmt.Scan(&amount)
	time.Sleep(time.Duration(amount) * time.Second)
	fmt.Println("time has passed now do your shish")
}
