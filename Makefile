marvin: *.go
	go build .

.PHONY: serve
serve: marvin
	./marvin

.PHONY: clean
clean:
	rm marvin
