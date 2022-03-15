package main

import (
	"errors"
	"flag"
	"github.com/streadway/amqp"
	"github.com/xh-dev-go/xhUtils/flagUtils"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
	"io/ioutil"
	"os"
)

const VERSION = "1.1.0"

var NoExchangeOrQueue = errors.New("exchange name and queue name can not be empty at the same time")

func validate(exchangeName, queueName string) error {
	if exchangeName == "" && queueName == "" {
		return NoExchangeOrQueue
	}
	return nil
}
func main() {
	url := flagString.New("amqp-url", "The connection string of amqp").BindCmd()
	exchangeName := flagString.NewDefault("exchange-name", "", "The name of the exchange").BindCmd()
	queueName := flagString.NewDefault("queue-name", "", "Tme name of the queue or routing key if exchange-name not empty").BindCmd()
	version := flagUtils.Version().BindCmd()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if err := validate(exchangeName.Value(), queueName.Value()); err != nil {
		panic(err)
	}

	if version.Value() {
		println(VERSION)
		os.Exit(0)
	}

	conn, err := amqp.Dial(url.Value())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	b, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	err = ch.Publish(exchangeName.Value(), queueName.Value(), false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        b,
	})
	if err != nil {
		panic(err)
	}
	println("done")
}
