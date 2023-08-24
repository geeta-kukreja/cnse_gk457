## Voter API - docker containerized

## Running with Docker

### Build and Run Using Single Docker Command:

Build and run the docker image:

```bash
docker compose up --build
```

### Use Published Docker Image for this project to run

If you prefer not to build the Docker image or have access to the source code, you can use the published Docker image directly. The image is available at [`geetakukreja/voter-api-container`](https://hub.docker.com/repository/docker/geetakukreja/voter-api-container/).

For API: http://localhost:1080  
Redis: http://localhost:8001

### Build and Run Using Docker Commands:

Build the docker image:

```bash
docker compose build
```

Run the docker container:

```bash
docker compose up


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