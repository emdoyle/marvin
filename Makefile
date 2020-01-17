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

.PHONY: build
build: marvin
	docker build . -t marvin:latest

.PHONY: tag
tag:
	docker tag marvin:latest emdoyle/marvin:latest

.PHONY: push
push:
	docker push emdoyle/marvin:latest

.PHONY: image
image:
	make build
	make tag
	make push
