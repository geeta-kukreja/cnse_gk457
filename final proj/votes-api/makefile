SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the vote api executable"
	@echo "	   run					Run the vote program from code"
	@echo "	   run-bin				Run the vote api executable"
	@echo "	   get-vote			Get a vote by id pass voteId=<id> on command line"
	@echo "	   get-all-votes				Get all votes information"
	@echo "    create-vote          	Add a vote record by passing voteId=<voteId> voterId=<voterId> pollId=<pollId> and voteValue=<voteValue> on command line"
	

.PHONY: build
build:
	go build .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./vote-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./vote-linux-arm64 .
	
.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./vote-api

.PHONY: get-vote
get-vote:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1082/votes/$(voteId) 

.PHONY: get-all-votes
get-all-votes:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1082/votes 

.PHONY: create-vote
create-vote:
	curl -w "\nHTTP Status: %{http_code}\n" -d '{ "voteId": $(voteId) ,"voterId": $(voterId) ,"pollId": $(pollId) , "voteValue": $(voteValue)}' -H "Content-Type: application/json" -X POST http://localhost:1082/votes

