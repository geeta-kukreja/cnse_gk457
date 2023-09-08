package main

import (
	"flag"
	"fmt"
	"poll-api/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	hostFlag string
	portFlag uint
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
	flag.UintVar(&portFlag, "p", 1080, "Default Port")

	flag.Parse()
}

// main is the entry point for our todo API application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {
	processCmdLineFlags()
	r := gin.Default()
	r.Use(cors.Default())

	apiHandler := api.NewPollApi()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	r.GET("/polls", apiHandler.ListAllPolls)
	r.POST("/polls", apiHandler.AddPoll)
	r.POST("/polls/:pollId/options/:pollOptionID", apiHandler.AddpollOption)

	r.GET("/polls/:pollId", apiHandler.GetPollById)
	r.GET("/polls/:pollId/options", apiHandler.GetAllPollOptions)
	r.GET("/voters/:pollId/options/:pollOptionID", apiHandler.GetPollsById)

	// r.PUT("/todo", apiHandler.UpdateToDo)
	// r.DELETE("/todo", apiHandler.DeleteAllToDo)
	// r.DELETE("/todo/:id", apiHandler.DeleteToDo)
	// r.GET("/crash", apiHandler.CrashSim)
	// r.GET("/health", apiHandler.HealthCheck)

	//We will now show a common way to version an API and add a new
	//version of an API handler under /v2.  This new API will support
	//a path parameter to search for todos based on a status
	// v2 := r.Group("/v2")
	// v2.GET("/todo", apiHandler.ListSelectTodos)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	r.Run(serverPath)
}

// func main() {
// 	//v := voter.NewVoter(1, "John", "Doe")
// 	//v.AddPoll(1)
// 	//v.AddPoll(2)
// 	//v.AddPoll(3)
// 	//v.AddPoll(4)

// 	//b, _ := json.Marshal(v)
// 	//fmt.Println(string(b))
// 	vl := api.NewVoterApi()

// 	//POST /voter with body of John Doe
// 	vl.AddVoter(1, "John", "Doe")

// 	//POST /voter/1/poll with body of a voter record only including the Poll information
// 	vl.LetsSimulateAPostForAPoll(1)
// 	fmt.Println()
// 	vl.LetsSimulateAPostForAPoll(5)

// 	/*
// 		vl.AddVoter(1, "John", "Doe")
// 		vl.AddPoll(1, 1)
// 		vl.AddPoll(1, 2)
// 		vl.AddVoter(2, "Jane", "Doe")
// 		vl.AddPoll(2, 1)
// 		vl.AddPoll(2, 2)

// 		fmt.Println("------------------------")
// 		fmt.Println(vl.GetVoterJson(1))
// 		fmt.Println("------------------------")
// 		fmt.Println(vl.GetVoterJson(2))
// 		fmt.Println("------------------------")
// 		fmt.Println(vl.GetVoterListJson())
// 		fmt.Println("------------------------")
// 	*/
// }
