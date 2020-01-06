marvin: src/** assets/build/**
	go build -o marvin ./src

assets/build/**: assets/src/**
	cd assets/ && yarn build

.PHONY: serve
serve: marvin
	./marvin

.PHONY: clean
clean:
	rm marvin
	rm -r assets/build/
