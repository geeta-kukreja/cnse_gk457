package voter

import (
	"encoding/json"
	"errors"
	"time"
)

type VoterPoll struct {
	PollID   uint
	VoteDate time.Time
}

type Voter struct {
	VoterID     uint
	FirstName   string
	LastName    string
	VoteHistory []VoterPoll
}
type VoterList struct {
	Voters map[uint]Voter //A map of VoterIDs as keys and Voter structs as values
}

// constructor for VoterList struct
func NewVoter(id uint, fn, ln string) *Voter {
	return &Voter{
		VoterID:     id,
		FirstName:   fn,
		LastName:    ln,
		VoteHistory: []VoterPoll{},
	}
}

func (v *VoterList) AddVoter(voter Voter) error {
	if _, exists := v.Voters[voter.VoterID]; exists {
		return errors.New("voter already exists")
	}

	v.Voters[voter.VoterID] = voter

	return nil
}
func (vl *VoterList) AddVoterPoll(voterID, pollID uint) (VoterPoll, error) {
	voter, exists := vl.Voters[voterID]
	if !exists {
		return VoterPoll{}, errors.New("voter does not exist")
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollID == pollID {
			return VoterPoll{}, errors.New("voter has already voted in this poll")
		}
	}

	newVoterPoll := VoterPoll{
		PollID:   pollID,
		VoteDate: time.Now(),
	}

	voter.VoteHistory = append(voter.VoteHistory, newVoterPoll)
	vl.Voters[voter.VoterID] = voter

	return newVoterPoll, nil
}

// func (v *Voter) AddPoll(voterId, pollID uint) error {
// 	// Create a new VoterPoll struct with the pollID and the current time
// 	newPoll := VoterPoll{
// 		PollID:   pollID,
// 		VoteDate: time.Now(),
// 	}

// 	for _, poll := range v.VoteHistory {
// 		if poll.PollID == pollID {
// 			return errors.New("voter has already voted in this poll")
// 		}
// 	}
// 	v.VoteHistory = append(v.VoteHistory, newPoll)

// 	// v.Voters[v.VoterID] = v
// 	return nil
// }

func (v *Voter) ToJson() string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (t *VoterList) GetItem(id uint) (Voter, error) {
	// Check if the item exists before trying to get it
	// This is a good practice, return an error if the
	// item does not exist
	item, ok := t.Voters[id]
	if !ok {
		return Voter{}, errors.New("Voter not found")
	}
	return item, nil
}

func (t *VoterList) GetAllPollsByVoterID(id uint) ([]VoterPoll, error) {
	voter, ok := t.Voters[id]
	if !ok {
		return nil, errors.New("Voter not found")
	}
	return voter.VoteHistory, nil
}

func (t *VoterList) GetAllPollsByVoterIDAndPollId(id uint, pollId uint) (VoterPoll, error) {
	voter, ok := t.Voters[id]
	if !ok {
		return VoterPoll{}, errors.New("Voter not found")
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollID == pollId {
			return poll, nil
		}
	}

	return VoterPoll{}, nil
}
