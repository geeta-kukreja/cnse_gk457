# cnse_gk457
Repository for CNSE Class - CLI projects
# 1. Voting API Project (final project)

The Voting API Project is a set of three interconnected APIs for managing and recording votes in polls. This README provides an overview of the project, how to run the APIs, and how to test them using a provided shell script.

## APIs Overview

The Voting API Project consists of three APIs:

1. **Votes API:** Manages votes, voters, and polls.
2. **Voters API:** Manages voter information.
3. **Polls API:** Manages polls and poll options.

These APIs work together to allow users to vote in polls and record their votes.

## Running the APIs

Before running the APIs, make sure you have the following prerequisites:

- Go (Golang) installed on your system.
- Redis server running (for caching votes).

Follow these steps to run the APIs:

1. Clone the project repository to your local machine and navigate to `voting-application` directory:

   ```bash
   git clone https://github.com/geeta-kukreja/cnse_gk457.git
   cd voting-application
   ```

## Running with Docker

### Method 1. Build and Run Using Single Docker Command:

Build and run the docker image:

```bash
docker compose up --build
```

-----

### Method 2. Build and Run Using Docker Commands:

Build the docker image:

```bash
docker compose build
```

Run the docker container:

```bash
docker compose up
```

-----

## Configuring Redis

The Voter API is designed to work seamlessly with Redis, which is set up to automatically start in its own Docker container before the Voter API starts up (as defined in `docker-compose.yml`). The two containers share the same network, making communication easy.

By default, the Voter API uses the URL `redis:6379` to establish a connection with the Redis container. This URL points to the Redis service within the Docker network. If you wish to use a different Redis instance or have a specific Redis server you'd like to connect to, you can configure this by setting the `REDIS_URL` environment variable for the Voter API container

## Testing the APIs

To test the APIs, a shell script (test-apis.sh) is provided. This script covers various scenarios for each API, including listing votes, retrieving votes by ID, adding votes, modifying votes, and deleting votes.

Follow these steps to test the APIs using the provided script:

1. Make the script executable:
  ```bash
   chmod +x test.sh
   ```

2. Run the script:
  ```bash
   ./test.sh
   ```

The script will execute different API requests and display the results for each scenario.


##  2. ToDo App


This is the class assignment for the ToDo list CLI assignment to implement a go language CLI tool to manage a list of todo items..

This application is driven by a simple text file based database.  


```
todo git:(main) ✗ ./todo -h
Usage of ./todo:
  -a string
        Add an item to the database
  -d int
        Delete an item from the database
  -db string
        Name of the database file (default "./data/todo.json")
  -l    List all the items in the database
  -q int
        Query an item in the database
  -s    Change item 'done' status to true or false
  -u string
        Update an item in the database
  ```


