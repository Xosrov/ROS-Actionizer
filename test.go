package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/bluenviron/goroslib/v2"
)

var node *goroslib.Node

type message struct {
	Msg string
}

func subscriber(ctx context.Context) {
	// create a subscriber
	sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:  node,
		Topic: "test/topic",
		Callback: func(msg *message) {
			fmt.Printf("Incoming: %+v\n", msg)
		},
	})
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	<-ctx.Done()

}

func publisher(ctx context.Context) {
	// create a publisher
	pub, err := goroslib.NewPublisher(goroslib.PublisherConf{
		Node:  node,
		Topic: "test/topic",
		Msg:   new(message),
	})
	if err != nil {
		panic(err)
	}
	defer pub.Close()

	r := node.TimeRate(1 * time.Second)

	for {
		select {
		// publish a message every second
		case <-r.SleepChan():
			msg := &message{
				Msg: "some message",
			}
			fmt.Printf("Outgoing: %+v\n", msg)
			pub.Write(msg)

		// handle CTRL-C
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	var err error
	node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          "goroslib_test",
		MasterAddress: "localhost:11311",
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go publisher(ctx)
	go subscriber(ctx)

	<-c
	cancel()
}
