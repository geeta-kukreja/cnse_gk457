package api

import (
	// ... other imports ...

	"fmt"
	"log"
	"net/http"
	"poll-api/poll"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PollApi struct {
	pollList *poll.PollCache
}

func NewPollApi() *PollApi {
	pollCache, _ := poll.NewPollCache()
	return &PollApi{
		pollList: pollCache,
	}
}

// Define API functions for managing polls just like you did for voters
// For example:
func (p *PollApi) ListAllPolls(c *gin.Context) {
	allPolls, err := p.pollList.GetAllPolls()
	if err != nil {
		log.Println("Error", err)
		return
	}
	c.JSON(http.StatusOK, allPolls)
}
func (v *PollApi) AddPoll(c *gin.Context) {
	// Parse the JSON request body to extract voter details
	var newPoll poll.Poll
	if err := c.BindJSON(&newPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newPoll = *poll.NewPoll(newPoll.PollID, newPoll.PollTitle, newPoll.PollQuestion)
	if err := v.pollList.AddPoll(newPoll); err != nil {
		log.Println("Error adding Poll: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// Respond with the newly created voter's ID
	c.JSON(http.StatusOK, newPoll)
}

func (va *PollApi) AddpollOption(c *gin.Context) {
	pollID := c.Param("pollId")
	pollIDUint, err := strconv.ParseUint(pollID, 10, 32)
	if err != nil {
		log.Println("Error converting poll ID to uint: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	optionID := c.Param("pollOptionID")

	optionIDUint, err := strconv.ParseUint(optionID, 10, 32)
	if err != nil {
		log.Println("Error converting poll option ID to uint: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var newpo poll.PollOption
	if err := c.BindJSON(&newpo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// optionValue := c.Param("pollOptionValue")
	newPollOption1 := *poll.NewpollOption(newpo.PollOptionID, newpo.PollOptionValue)
	newpollOption, err := va.pollList.AddpollOption(uint(pollIDUint), uint(optionIDUint), newPollOption1)
	if err != nil {
		log.Println("Error adding  poll option: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Printf("newPollOption1: %+v\n", newPollOption1)
	fmt.Printf("newPollOption1: %+v\n", newpollOption)
	c.JSON(http.StatusOK, newpollOption)
}

func (v *PollApi) GetPollById(c *gin.Context) {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id := c.Param("pollId")
	id64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Println("Error converting id to int64: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	voter, err := v.pollList.GetItem(uint(id64))
	if err != nil {
		log.Println("Item not found: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	c.JSON(http.StatusOK, voter)
}

func (v *PollApi) GetAllPollOptions(c *gin.Context) {
	pollID := c.Param("pollId")
	id64, err := strconv.ParseInt(pollID, 10, 32)
	polls, err := v.pollList.GetAllPollOptions(uint(id64))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, poll := range polls {
			fmt.Println("Poll Option ID:", poll.PollOptionID)
			fmt.Println("Poll Option Value:", poll.PollOptionValue)
		}
	}
	c.JSON(http.StatusOK, polls)

}

func (v *PollApi) GetPollsById(c *gin.Context) {
	pollID := c.Param("pollId")
	optionID := c.Param("pollOptionId")
	id64, err := strconv.ParseInt(pollID, 10, 32)
	pollId64, err := strconv.ParseInt(optionID, 10, 32)
	polls, err := v.pollList.GetPollOptionsByID(uint(id64), uint(pollId64))
	if err != nil {
		fmt.Println("Error:", err)
	}
	c.JSON(http.StatusOK, polls)

}
