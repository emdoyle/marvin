marvin: src
	go build -o marvin ./src

.PHONY: serve
serve: marvin
	./marvin

.PHONY: clean
clean:
	rm marvin
