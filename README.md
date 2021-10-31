# go-simple-payment
i'am inspired from this article [money-scale-strong-data](https://eng.uber.com/money-scale-strong-data/), and try to make simple implementation of order domain using go.

## how to run
```bash
docker run --name postgres-sadbox --rm -d \
  -e POSTGRES_PASSWORD=sandbox \
  -e POSTGRES_USER=sandbox \
  -e POSTGRES_DB=sandbox \
  -p 6543:5432 \
  postgres:13-alpine

make test
make
```

## how to use 
```sh
curl -H 'Content-Type: application/json' \
-d '{
    "entries": [
        { "name": "Trip fare", "amount": -18, "userID": "1" },
        { "name": "Service fee", "amount": -2, "userID": "1" },
        { "name": "Trip fare", "amount": 18, "userID": "2" },
        { "name": "Service fee", "amount": 2, "userID": "3" }
    ]
}' localhost:3000/v0/order

curl -H 'Content-Type: application/json' \
-d '{
    "entries": [
        { "name": "Trip fare", "amount": -18, "userID": "1" },
        { "name": "Service fee", "amount": -2, "userID": "1" },
        { "name": "Trip fare", "amount": 18, "userID": "2" },
        { "name": "Service fee", "amount": 3, "userID": "3" }
    ]
}' localhost:3000/v0/order
```


