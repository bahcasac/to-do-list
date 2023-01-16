include .env

uname := $(shell uname && uname -p)
ifeq ($(uname), Darwin arm)
	run_cmd := dynamic
else
	run_cmd := musl
endif

GO_RELOAD = $(GOPATH)/bin/reflex -s -r '\.go$$' --

up-postgres:
	docker compose up -d postgres

deps:
	go mod tidy && go get -u -tags $(run_cmd) ./cmd/api

run-api:
	$(GO_RELOAD) go run -tags $(run_cmd) -race ./cmd/api/main.go run-api