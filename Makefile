run:
	go run app/main.go

test:
	gotest -v ./... -covermode=atomic
