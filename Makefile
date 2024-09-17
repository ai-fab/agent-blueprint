.PHONY: build run test clean

build:
	go build -o service-blueprint

run:
	go run . serve

test:
	go test ./...

clean:
	rm -f service-blueprint

clean-db:
	rm -rf pb_data

migrate: build
	./service-blueprint migrate up

migrate-down: build
	./service-blueprint migrate down

migrate-fresh: build
	./service-blueprint migrate down
	./service-blueprint migrate up
