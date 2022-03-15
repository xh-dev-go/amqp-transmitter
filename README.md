# amqp-transmitter

## install
```
go install github.com/xh-dev-go/amqp-transmitter
```

## build
```
go build
```

## Command
```shell
$ ./amqp-transmitter
  -amqp-url string
        The connection string of amqp
  -exchange-name string
        The name of the exchange
  -queue-name string
        Tme name of the queue or routing key if exchange-name not empty
  -version
        show application version
# Demo
./amqp-transmitter --amqp-url {url} --exchange-name {exchange-name} --queue-name {queue name}
```

