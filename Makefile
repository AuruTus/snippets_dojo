# TODO add init for tstr_entry

.PHONY: build
build:
	go build -o ./bin/app .

.PHONY: run
run: build
	./bin/app
