SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the voter api executable"
	@echo "	   run					Run the voter program from code"
	@echo "	   run-bin				Run the voter executable"
	@echo "	   load-db				Add sample data via curl"
	@echo "	   get-by-id			Get a voter by id pass id=<id> on command line"
	@echo "	   get-all				Get all voters information"
	@echo "    add-voter          	Add a voter record by passing id=<voterId> firstName=<firstName> & lastName=<lastName> on command line"
	@echo "    get-voter-history  	Fetch voter history for the voter with id=<:voterId>"
	@echo "    get-poll     		Fetch single voter poll data with id=<:voterId> and pollid=<:pollId>"
	@echo "    add-poll-by-id     		Add a voter poll record for the voter with id=<:voterId> and pollid=<:pollId>"
	@echo "    health-check       	Check about Voter's api"

.PHONY: build
build:
	go build .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./voter-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./voter-linux-arm64 .
	
.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./voters-api


.PHONY: load-db
load-db:
	curl -d '{ "VoterID":0,"FirstName":"John","LastName":"Doe","VoteHistory":[]}' -H "Content-Type: application/json" -X POST http://localhost:1081/voters 
	curl -d '{ "VoterID":1,"FirstName":"GEETA","LastName":"KUKREJA","VoteHistory":[]}' -H "Content-Type: application/json" -X POST http://localhost:1081/voters


.PHONY: get-by-id
get-by-id:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voters/$(id) 

.PHONY: update-2
update-2:
	curl -d '{ "id": 2, "title": "$(title)", "done": false }' -H "Content-Type: application/json" -X PUT http://localhost:1081/todo 


.PHONY: get-all
get-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voters 

.PHONY: get-voter-history
get-voter-history:
	curl -w "\nHTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voters/$(id)/polls

.PHONY: get-poll
get-poll:
	curl -w "\nHTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voters/$(id)/polls/$(pollid)
.PHONY: add-voter
add-voter:
	curl -w "\nHTTP Status: %{http_code}\n" -d '{ "VoterID":$(id), "firstName": "$(firstName)", "lastName": "$(lastName)" }' -H "Content-Type: application/json" -X POST http://localhost:1081/voters

.PHONY: add-poll-id
add-poll-id:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X POST http://localhost:1081/voters/$(id)/polls/$(pollid) 

.PHONY: health-check
health-check:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/voters/health

.PHONY: delete-all
delete-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1081/todo 

.PHONY: delete-by-id
delete-by-id:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1081/todo/$(id) 

.PHONY: get-v2
get-v2:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/v2/todo?done=$(done) 

.PHONY: get-v2-all
get-v2-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1081/v2/todo
