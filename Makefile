
.PHONY: init
init:
ifeq (,$(wildcard ./tstr_entry.go))
	./generator.sh
endif

.PHONY: build
build: init
	go build -o ./bin/app .

.PHONY: run
run: build
	./bin/app
