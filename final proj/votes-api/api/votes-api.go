package api

import (
	// ... other imports ...

	"log"
	"net/http"
	"strconv"
	election "votes-api/votes"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type VoteApi struct {
	voteList    *election.VoteCache
	pollAPIURL  string
	voterAPIURL string
	apiClient   *resty.Client
}

func NewVoteApi(pollAPIURL string, voterAPIURL string) *VoteApi {
	voteCache, _ := election.NewVoteCache()
	apiClient := resty.New()
	return &VoteApi{
		voteList:    voteCache,
		pollAPIURL:  pollAPIURL,
		voterAPIURL: voterAPIURL,
		apiClient:   apiClient,
	}
}

// Define API functions for managing polls just like you did for voters
// For example:
func (p *VoteApi) ListAllVotes(c *gin.Context) {
	allVotes, err := p.voteList.GetAllVotes()
	if err != nil {
		log.Println("Error", err)
		return
	}
	c.JSON(http.StatusOK, allVotes)
}
func (v *VoteApi) AddVote(c *gin.Context) {
	// Parse the JSON request body to extract voter details
	var newVote election.Vote
	if err := c.BindJSON(&newVote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newVote = *election.NewVote(newVote.VoteID, newVote.VoterID, newVote.PollID, newVote.VoteValue)
	if err := v.voteList.AddVote(newVote); err != nil {
		log.Println("Error adding Vote: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// Respond with the newly created voter's ID
	c.JSON(http.StatusOK, newVote)
}

func (v *VoteApi) GetVoteById(c *gin.Context) {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id := c.Param("voteId")
	id64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Println("Error converting id to int64: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	vote, err := v.voteList.GetItem(uint(id64))
	if err != nil {
		log.Println("Item not found: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	c.JSON(http.StatusOK, vote)
}
