package main

import "go-interview/consul"

func main() {
	register_client := consul.NewRegistryClient("192.168.1.8", 12456)
	register_client.Register()
}
