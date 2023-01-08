
.PHONY: init
init:
# if the project is new cloned from remote repo, the tstr_entry.go will be missing
ifeq (,$(wildcard ./tstr_entry.go))
	./scripts/init_dojo.sh
endif

.PHONY: init_tstr
# init the new tester and its go file under ./src/test_snippets
init_tstr:
	./scripts/init_test_snippet.sh $(file) $(tstr)

.PHONY: build
build: init
	go build -o ./bin/app .

.PHONY: run
run: build
	./bin/app
