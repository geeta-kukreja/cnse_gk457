#!/bin/bash

# Define your API endpoints
voter_api_url="http://localhost:1081"
poll_api_url="http://localhost:1080"
votes_api_url="http://localhost:1082"

# Test Voter APIS Scenarios
echo "Testing Voter API Scenarios..."

# Scenario 1: POST request to add voter
echo "Scenario 1: Add voters"
curl -d '{ "id":1, "firstName": "Geeta", "lastName": "Kukreja" }' -H "Content-Type: application/json" -X POST $voter_api_url/voters
curl -d '{ "id":0,"FirstName":"John","LastName":"Doe","VoteHistory":[]}' -H "Content-Type: application/json" -X POST $voter_api_url/voters

# Scenario 3: GET request to list all voters
echo "Scenario 3: List all voters"
curl $headers -X GET $voter_api_url/voters

# Scenario 4: GET request to retrieve a specific voter
echo "Scenario 4: Retrieve item by ID"
curl $headers -X GET $voter_api_url/voters/1

# Test API 2 Scenarios
echo "Testing Poll API Scenarios..."

# Scenario 5: POST request to create a new item
echo "Scenario 5: Create a new poll"
curl -d '{ "pollId": 1, "title": "Cusine", "question": "What is your Favourite Cuisine"}' -H "Content-Type: application/json" -X POST $poll_api_url/polls
curl -d '{ "pollId": 2, "title": "Favourite City", "question": "Your favourite sport?"}' -H "Content-Type: application/json" -X POST $poll_api_url/polls

# Scenario 6: POST options to poll
echo "Scenario 6: Add options to poll"
curl -d '{ "pollOptionID":1, "pollOptionValue": "Indian" }' -H "Content-Type: application/json" -X POST $poll_api_url/polls/1/options/1
curl -d '{ "pollOptionID":2,"pollOptionValue": "Chinese" }' -H "Content-Type: application/json" -X POST $poll_api_url/polls/1/options/2
curl -d '{ "pollOptionID":1, "pollOptionValue": "Paris" }' -H "Content-Type: application/json" -X POST $poll_api_url/polls/2/options/1
curl -d '{ "pollOptionID":2, "pollOptionValue": "New York" }' -H "Content-Type: application/json" -X POST $poll_api_url/polls/2/options/2
curl -d '{ "pollOptionID":3, "pollOptionValue": "Delhi" }' -H "Content-Type: application/json" -X POST $poll_api_url/polls/2/options/3

# Scenario 8: Get poll
echo "Scenario 8: Retrieve poll by id"
curl $headers -X GET $poll_api_url/polls/1
curl $headers -X GET $poll_api_url/polls/2

# Scenario 9: Get all poll
echo "Scenario 9: Retrieve all polls"
curl $headers -X GET $poll_api_url/polls

# Scenario 10: Get poll options
echo "Scenario 10: Get poll options by id"
curl $headers -X GET $poll_api_url/polls/1/options

# Test Votes API Scenarios
echo "Testing Votes API Scenarios..."

# Scenario 11: PUT request to modify an item by ID (replace ID)
echo "Scenario 11: Add vote in votes list and update voter history in voter"
curl -d '{ "voteId": 1, voterId": 1, "pollId": 1, "voteValue": 1 }' -H "Content-Type: application/json" -X POST $votes_api_url/votes/1
curl -d '{ "voteId": 2, "voterId": 2, "pollId": 1, "voteValue": 1 }' -H "Content-Type: application/json" -X POST $votes_api_url/votes/2
curl -d '{ "voteId": 3, "voterId": 1, "pollId": 2, "voteValue": 2 }' -H "Content-Type: application/json" -X POST $votes_api_url/votes/3
curl -d '{ "voteId": 4, "voterId": 2, "pollId": 2, "voteValue": 3 }' -H "Content-Type: application/json" -X POST $votes_api_url/votes/4

# Scenario 12: Get all votes.
echo "Scenario 12: Retrieve all votes"
curl $headers -X GET $votes_api_url/votes

# Scenario 12: Get specific vote.
echo "Scenario 12: Get specific vote by id"
curl $headers -X GET $votes_api_url/votes/1

# End of testing scenarios
echo "Testing complete."
