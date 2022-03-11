package main

import (
	"flag"
	"github.com/streadway/amqp"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
	"io/ioutil"
	"os"
)

func main() {
	url := flagString.New("amqp-url", "The connection string of amqp").BindCmd()
	queueName := flagString.New("queue-name", "The name of the queue").BindCmd()

	flag.Parse()
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
	err = ch.Publish("", queueName.Value(), false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        b,
	})
	if err != nil {
		panic(err)
	}
	println("done")
}
