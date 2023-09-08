package main

import (
	"flag"
	"fmt"
	"votes-api/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	hostFlag    string
	portFlag    uint
	voterAPIURL string
	pollAPIURL  string
)

func processCmdLineFlags() {

	//Note some networking lingo, some frameworks start the server on localhost
	//this is a local-only interface and is fine for testing but its not accessible
	//from other machines.  To make the server accessible from other machines, we
	//need to listen on an interface, that could be an IP address, but modern
	//cloud servers may have multiple network interfaces for scale.  With TCP/IP
	//the address 0.0.0.0 instructs the network stack to listen on all interfaces
	//We set this up as a flag so that we can overwrite it on the command line if
	//needed
	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.StringVar(&voterAPIURL, "v", "http://host.docker.internal:1081", "Default voter API location")
	flag.StringVar(&pollAPIURL, "papi", "http://host.docker.internal:1080", "Default poll API location")
	flag.UintVar(&portFlag, "p", 1082, "Default Port")

	flag.Parse()
}

// main is the entry point for our todo API application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {
	processCmdLineFlags()
	r := gin.Default()
	r.Use(cors.Default())

	apiHandler := api.NewVoteApi(voterAPIURL, pollAPIURL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	r.GET("/votes", apiHandler.ListAllVotes)
	r.POST("/votes", apiHandler.AddVote)

	r.GET("/votes/:voteId", apiHandler.GetVoteById)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	r.Run(serverPath)
}
