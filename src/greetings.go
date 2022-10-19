package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Hey, nice to meet you. My name is %s!\n", PickName())
}

func PickName() string {
	var nameList [3]string
	nameList[0] = "John"
	nameList[1] = "Hrant"
	nameList[2] = "Anna" 
	rand.Seed(time.Now().UnixNano())
	return nameList[rand.Intn(len(nameList))]
}


