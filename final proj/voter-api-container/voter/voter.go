package voter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type VoterPoll struct {
	PollID   uint      `json:"pollId"`
	VoteDate time.Time `json:"voteDate"`
}

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "voter:"
)

type Voter struct {
	VoterID     uint        `json:"voterId"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	VoteHistory []VoterPoll `json:"voteHistory"`
}

//	type VoterList struct {
//		Voters map[uint]Voter //A map of VoterIDs as keys and Voter structs as values
//	}
type cache struct {
	cacheClient *redis.Client
	jsonHelper  *rejson.Handler
	context     context.Context
}

type VoterList struct {
	Voters map[uint]Voter
}
type VoterCache struct {
	cache
}

func NewVoterCache() (*VoterCache, error) {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}

	return NewWithCacheInstance(redisUrl)
}

// The constructor function that returns a pointer to a new VoterCache.
func NewWithCacheInstance(location string) (*VoterCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	ctx := context.Background()

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error())
		return nil, err
	}

	jsonHelper := rejson.NewReJSONHandler()
	jsonHelper.SetGoRedisClientWithContext(ctx, client)

	return &VoterCache{
		cache: cache{
			cacheClient: client,
			jsonHelper:  jsonHelper,
			context:     ctx,
		},
	}, nil
}
func redisKeyFromId(id uint) string {
	return fmt.Sprintf("%s%d", RedisKeyPrefix, id)
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
func (vc *VoterCache) getItemFromRedis(key string, item *Voter) error {
	itemObject, err := vc.jsonHelper.JSONGet(key, ".")
	if err != nil {
		return err
	}

	err = json.Unmarshal(itemObject.([]byte), item)
	if err != nil {
		return err
	}

	return nil
}
func (v *VoterCache) AddVoter(voter Voter) error {
	if _, err := v.GetItem(voter.VoterID); err == nil {
		return errors.New("voter already exists")
	}

	// v.Voters[voter.VoterID] = voter
	redisKey := redisKeyFromId(voter.VoterID)
	if _, setErr := v.jsonHelper.JSONSet(redisKey, ".", voter); setErr != nil {
		return setErr
	}
	return nil
}
func (v *VoterCache) AddVoterPoll(voterID, pollID uint) (VoterPoll, error) {
	voter, err := v.GetItem(voterID)
	if err != nil {
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

	redisKey := redisKeyFromId(voter.VoterID)

	if _, setErr := v.jsonHelper.JSONSet(redisKey, ".", voter); setErr != nil {
		return VoterPoll{}, setErr
	}

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

func (v *VoterCache) GetItem(id uint) (Voter, error) {
	// Check if the item exists before trying to get it
	// This is a good practice, return an error if the
	// item does not exist
	var voter Voter

	redisKey := redisKeyFromId(id)
	if err := v.getItemFromRedis(redisKey, &voter); err != nil {
		return Voter{}, errors.New("voter does not exist")
	}

	return voter, nil
}
func (v *VoterCache) GetAllVoters() ([]Voter, error) {
	var voters []Voter

	pattern := fmt.Sprintf("%s*", RedisKeyPrefix)
	keys, err := v.cacheClient.Keys(v.context, pattern).Result()

	if err != nil {
		return voters, err
	}

	for _, key := range keys {
		var voter Voter

		err := v.getItemFromRedis(key, &voter)
		if err != nil {
			return voters, err
		}

		voters = append(voters, voter)
	}

	return voters, nil
}
func (v *VoterCache) GetAllPollsByVoterID(id uint) ([]VoterPoll, error) {
	voter, err := v.GetItem(id)
	if err != nil {
		return nil, errors.New("Voter not found")
	}
	return voter.VoteHistory, nil
}

func (v *VoterCache) GetPollByVoterIDAndPollId(id uint, pollId uint) (VoterPoll, error) {
	voter, err := v.GetItem(id)
	if err != nil {
		return VoterPoll{}, errors.New("Voter not found")
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollID == pollId {
			return poll, nil
		}
	}

	return VoterPoll{}, nil
}
