package main

import (
	route "Challenge/internal/adapters/api"
	kafka "Challenge/internal/repositories/transactions"
)

func main() {
	go kafka.Consume()
	route.HandleRequest()
}
