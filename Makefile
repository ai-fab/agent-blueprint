.PHONY: build run test clean

build:
	go build -o service-blueprint

run: build
	./service-blueprint serve

test:
	go test ./...

clean:
	rm -f service-blueprint

migrate:
	./service-blueprint migrate up

migrate-down:
	./service-blueprint migrate down

migrate-fresh:
	./service-blueprint migrate down
	./service-blueprint migrate up
