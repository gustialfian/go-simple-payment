run:
	go run cmd/orderinsertionsvc/main.go \
  		--db "host=localhost port=6543 dbname=sandbox user=sandbox password=sandbox sslmode=disable" \
  		--port 8000
test:
	GO111MODULE=on go test ./...