test:
	go test ./... -timeout 10s --count 1 --race

coverage:
	go test ./... -cover