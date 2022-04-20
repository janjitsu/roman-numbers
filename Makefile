build:
	docker image build -t roman-numbers .

container:
	docker container run -p 8080:8080 roman-numbers

run:
	go run app/http/server.go

test:
	gotest -v ./... -covermode=atomic
