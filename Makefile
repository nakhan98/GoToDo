.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build
build:
	go build  -o gotodo ./cmd/gotodo

build-verbose:
	go build -a -x -o gotodo ./cmd/gotodo

clean:
	find ./ -name "*~" | xargs rm
	rm -f gotodo
