
.PHONY: all clean

all: onion.go 
	@go build
	./onion migratedb
	./onion serve

cleandb:
	rm -r *.db

migrate:
	./onion migratedb

run:
	./onion serve

