all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/jack/jack.go

check-cli: build
	./jack examples/create_database.sql
	./jack examples/create_table.sql
