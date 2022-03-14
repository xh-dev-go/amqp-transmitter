# amqp-transmitter

## install
```
go install github.com/xh-dev-go/amqp-transmitter@latest
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
  -queue-name string
        The name of the queue
  -version
        show application version

# Demo
echo "[message]$(date)" | ./amqp-transmitter --amqp-url {url} --queue-name {queue name}
```

