# RabbitMQ Simple

## Prerequisites
- [Go] - Golang
- [Docker] - Docker


[Go]: <https://go.dev/>
[Docker]: <https://www.docker.com/>

## How to start
**Step 1**  
Run docker compose file
```go sh
docker-compose up -d
```

**Step 2**  
Run setting file in target folder (Use command in their README)

**Step 3**  
Run consumer and producer at root project (Use command in the same README as step 2)

## How to monitor messages in RabbitMQ
http://localhost:15672/
```go sh
username : admin
password : 1234
( you can see username password in docker-compose file )
```

