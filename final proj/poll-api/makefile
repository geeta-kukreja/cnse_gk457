SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the poll api executable"
	@echo "	   run					Run the poll program from code"
	@echo "	   run-bin				Run the poll executable"
	@echo "	   load-sample-poll		Add sample data via curl"
	@echo "	   get-poll			Get a poll by id pass pollId=<id> on command line"
	@echo "	   get-all				Get all polls information"
	@echo "    add-poll          	Add a poll record by passing pollId=<pollId> title=<title> question=<question> on command line"
	@echo "    get-poll-options  	Fetch poll history for the poll with pollId=<:pollId>"
	@echo "    get-poll-option     		Fetch single poll poll data with pollId=<:pollId> and pollOptionID=<:pollOptionID>"
	@echo "    add-polloption     		Add a poll option record for the poll with pollId=<:pollId> pollOptionID=<:pollOptionID> and pollOptionValue=<pollOptionValue>"

.PHONY: build
build:
	go build .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./poll-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./poll-linux-arm64 .
	
.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./poll-api

.PHONY: load-sample-poll
load-sample-poll:
	curl -d '{ "pollId":0,"title":"Favorite Pet","question":"What type of pet do you like best?"}' -H "Content-Type: application/json" -X POST http://localhost:1080/polls 

.PHONY: get-poll
get-poll:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/polls/$(pollId) 

.PHONY: get-all
get-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/polls 

.PHONY: get-poll-options
get-poll-options:
	curl -w "\nHTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/polls/$(pollId)/options

.PHONY: get-poll-option
get-poll-option:
	curl -w "\nHTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/polls/$(pollId)/options/$(pollOptionID)
	
.PHONY: add-poll
add-poll:
	curl -w "\nHTTP Status: %{http_code}\n" -d '{ "pollId": $(pollId) , "title": "$(title)" , "question": "$(question)"}' -H "Content-Type: application/json" -X POST http://localhost:1080/polls

.PHONY: add-polloption
add-polloption:
	curl -w "HTTP Status: %{http_code}\n" -d '{ "pollId": $(pollId) , "pollOptionID":$(pollOptionID), "pollOptionValue": "$(pollOptionValue)"}' -H "Content-Type: application/json" -X POST http://localhost:1080/polls/$(pollId)/options/$(pollOptionID) 
