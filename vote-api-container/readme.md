## Voter API

Voter API for:
GET /voters - Get all voter resources including all voter history for each voter 
GET&POST /voters/:id - Get a single voter resource with voterID=:id including their entire voting history.  POST version adds one to the "database" -> Done

GET /voters/:id/polls - Gets the JUST the voter history for the voter with VoterID = :id  -> 

GET&POST /voters/:id/polls/:pollid - Gets JUST the single voter poll data with PollID = :id and 
VoterID = :id.  POST version adds one to the "database" -> Done
`
GET /voters/health - Returns a "health" record indicating that the voter API is functioning properly and some metadata about the API.  
Note the payload can be hard coded 

```
------------------------
```