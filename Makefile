.PHONY: build
build:
	go build -v -o bin/tx-bank cmd/*.go

.PHONY: build
run:
	@echo " >> build tx-bank"
	@make build
	@echo " >> tx-bank built."
	@echo " >> executing tx-bank"
	@./bin/tx-bank
	@echo " >> tx-bank is running"

.PHONY: docker-up
docker-up:
	docker-compose up --build -d

.PHONY: docker-down
docker-down:
	docker-compose down
