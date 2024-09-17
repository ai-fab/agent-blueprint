.PHONY: build run test clean clean-db migrate migrate-down migrate-fresh test-crud

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

test-crud: build
	./service-blueprint serve &
	sleep 5
	bash test_crud.sh
	pkill -f "./service-blueprint serve"
