![direct-exchange](../images/direct-exchange.png)

**Run Setting**
```go sh
QUEUE=asia.thailand.queue ROUTING_KEY=asia.thailand go run direct-exchange/setting.go
QUEUE=asia.singapore.queue ROUTING_KEY=asia.singapore go run direct-exchange/setting.go
```

**Run Consumer**
```go sh
QUEUE=asia.thailand.queue go run consumer.go
QUEUE=asia.singapore.queue go run consumer.go
```

**Run Producer**
```go sh
EXCHANGE_NAME=asia.exchange ROUTING_KEY=asia.thailand go run producer.go
EXCHANGE_NAME=asia.exchange ROUTING_KEY=asia.singapore go run producer.go
```

