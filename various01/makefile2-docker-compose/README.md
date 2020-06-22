## Influx DB 
Influx db port is 8086 - pod port is mapped to 8087(or other) of host

## Manually Create and Run "API Image ONLY"
```
docker build . -t podstatestore-api:big

docker run -d \
-p 8089:8081 \ 
podstatestore-api:big
#podstatestore-api:scratch 
#podstatestore-api:manual
#--net="host" \
#podstatestore-api:testing
```
## Docker Compose - Full Stack
To Run:
docker-compose up  


