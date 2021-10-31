https://eng.uber.com/money-scale-strong-data/


```bash
# shell 1
docker run --name postgres-sadbox --rm -d \
  -e POSTGRES_PASSWORD=sandbox \
  -e POSTGRES_USER=sandbox \
  -e POSTGRES_DB=sandbox \
  -p 6543:5432 \
  postgres:13-alpine

# shell 2
make
```

```sh
curl -v localhost:3000/v0/order

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


