.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build
build:
	mkdir build/
	go build  -o build/gotodo ./cmd/gotodo

build-verbose:
	go build -a -x -o gotodo ./cmd/gotodo

clean:
	find ./ -name "*~" | xargs rm
	rm -f gotodo
	rm -rf build
