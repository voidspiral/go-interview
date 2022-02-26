package main

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.1.9876"}),
		consumer.WithGroupName("mxshop"),
		)
	if err := c.Subscribe("test", consumer.MessageSelector{})
}
