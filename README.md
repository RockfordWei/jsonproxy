# JSON/XML Proxy Demo

A demo server to host a JSON/XML proxy in GO lang.

## Quick Start

The following script can clone and run a server on port 9090

``` bash
$ git clone https://github.com/RockfordWei/jsonproxy.git
$ cd jsonproxy && docker-compose build
$ docker-compose up -d
```

To stop this local server, just run `docker-compose down` in the same project folder.


## Sample GET

testing `GET /orders/{id}`

``` bash
$ curl "http://localhost:9090/orders/aeffb38f-a1a0-48e7-b7a8-2621a2678534"
```
which is expecting something like `{"id":"aeffb38f-a1a0-48e7-b7a8-2621a2678534","data":"mydataishere","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`


## Sample POST

testing `POST /orders`

``` bash
$ curl -X POST "http://localhost:9090/orders" -d '{"data":"some random test"}'
```

Output should look like `{"id":"13a7fc17-6221-4511-a935-39b53616d161","data":"some random test","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`


