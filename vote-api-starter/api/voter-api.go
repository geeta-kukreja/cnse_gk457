package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"voter-api-starter/voter"

	"github.com/gin-gonic/gin"
)

type VoterApi struct {
	voterList voter.VoterList
}

func NewVoterApi() *VoterApi {
	return &VoterApi{
		voterList: voter.VoterList{
			Voters: make(map[uint]voter.Voter),
		},
	}
}
func (v *VoterApi) ListAllVoters(c *gin.Context) {
	allVoters := v.voterList.Voters

	// Convert the map of voters to a slice of voters
	var voters []voter.Voter
	for _, v := range allVoters {
		voters = append(voters, v)
	}

	// Respond with the list of voters as JSON
	c.JSON(http.StatusOK, voters)

}

// func (v *VoterApi) AddVoter(voterID uint, firstName, lastName string) {
// 	v.voterList.Voters[voterID] = *voter.NewVoter(voterID, firstName, lastName)
// }

func (v *VoterApi) AddVoter(c *gin.Context) {
	// Parse the JSON request body to extract voter details
	var newVoter voter.Voter
	if err := c.BindJSON(&newVoter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newVoter = *voter.NewVoter(newVoter.VoterID, newVoter.FirstName, newVoter.LastName)
	if err := v.voterList.AddVoter(newVoter); err != nil {
		log.Println("Error adding voter: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// Respond with the newly created voter's ID
	c.JSON(http.StatusOK, gin.H{"voterID": newVoter.VoterID})
}

// func (v *VoterApi) AddPoll(c *gin.Context) {
// 	idS := c.Param("id")
// 	id64, err := strconv.ParseInt(idS, 10, 32)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	// Parse the JSON request body to extract voterID and pollID
// 	var requestData struct {
// 		VoterID uint `json:"voterID"`
// 		PollID  uint `json:"pollID"`
// 	}

// 	if err := c.BindJSON(&requestData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	// Check if the voterID exists in the voterList
// 	voter, exists := v.voterList.Voters[uint(id64)]

// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Voter not found"})
// 		return
// 	}

// 	// Add the pollID to the voter's polls
// 	err = voter.AddPoll(uint(id64), requestData.PollID)
// 	if err != nil {
// 		log.Println("Error adding voter poll: ", err)
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	// v.voterList.Voters[requestData.VoterID] = voter

// 	c.JSON(http.StatusOK, gin.H{"message": "Poll added successfully" + strconv.Itoa(int(requestData.PollID))})
// }

func (va *VoterApi) AddPoll(c *gin.Context) {
	voterID := c.Param("id")
	voterIDUint, err := strconv.ParseUint(voterID, 10, 32)
	if err != nil {
		log.Println("Error converting voter ID to uint: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pollID := c.Param("pollid")
	pollIDUint, err := strconv.ParseUint(pollID, 10, 32)
	if err != nil {
		log.Println("Error converting poll ID to uint: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newVoterPoll, err := va.voterList.AddVoterPoll(uint(voterIDUint), uint(pollIDUint))
	if err != nil {
		log.Println("Error adding voter poll: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newVoterPoll)
}

func (v *VoterApi) GetVoterById(c *gin.Context) {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Println("Error converting id to int64: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	voter, err := v.voterList.GetItem(uint(id64))
	if err != nil {
		log.Println("Item not found: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	c.JSON(http.StatusOK, voter)
}

func (v *VoterApi) GetPollsByVoterId(c *gin.Context) {

	// Assuming you have a `voterList` object of type `VoterList`
	voterID := c.Param("id") // Replace with the voter ID you want to retrieve polls for
	id64, err := strconv.ParseInt(voterID, 10, 32)
	polls, err := v.voterList.GetAllPollsByVoterID(uint(id64))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, poll := range polls {
			fmt.Println("Poll ID:", poll.PollID)
			fmt.Println("Vote Date:", poll.VoteDate)
		}
	}
	c.JSON(http.StatusOK, polls)

}

func (v *VoterApi) GetPollsById(c *gin.Context) {

	// Assuming you have a `voterList` object of type `VoterList`
	voterID := c.Param("id")    // Replace with the voter ID you want to retrieve polls for
	pollID := c.Param("pollid") // Replace with the voter ID you want to retrieve polls for
	id64, err := strconv.ParseInt(voterID, 10, 32)
	pollId64, err := strconv.ParseInt(pollID, 10, 32)
	polls, err := v.voterList.GetAllPollsByVoterIDAndPollId(uint(id64), uint(pollId64))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// for _, poll := range polls. {
		// 	fmt.Println("Poll ID:", poll.PollID)
		// 	fmt.Println("Vote Date:", poll.VoteDate)
		// }
	}
	c.JSON(http.StatusOK, polls)

}

func (v *VoterApi) GetVoterJson(voterID uint) string {
	voter := v.voterList.Voters[voterID]
	return voter.ToJson()
}

func (v *VoterApi) GetVoterList() voter.VoterList {
	return v.voterList
}

func (v *VoterApi) GetVoterListJson() string {
	b, _ := json.Marshal(v.voterList)
	return string(b)
}

func (va *VoterApi) HealthCheck(c *gin.Context) {
	healthRecord := map[string]interface{}{
		"status":  "healthy",
		"version": "1.0.0",
		"message": "Voter API is functioning properly",
	}
	c.JSON(200, healthRecord)
}
