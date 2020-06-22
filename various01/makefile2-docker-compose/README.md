## Influx DB 
Influx db port is 8086 - pod port is mapped to 8087(or other) of host

## With only influxdb part in docker-compose.yml

To Run:
docker-compose up  

#### docker-compose.yml:  

```
version: "3"

services:

  # influx container with database volume mounted at $PWD/influxdbstore
  influxdbstore:
    image: "influxdb"
    ports:
      - "8087:8086"
    networks:
      - shared
    volumes:
      - "./influxdbstore:/var/lib/influxdb"
    environment:
      - INFLUXDB_DB=metrics

networks:
  shared:
```
