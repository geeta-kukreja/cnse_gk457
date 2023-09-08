package poll

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type PollOption struct {
	PollOptionID    uint   `json:"pollOptionID"`
	PollOptionValue string `json:"pollOptionValue"`
}

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "poll:"
)

type Poll struct {
	PollID       uint         `json:"pollId"`
	PollTitle    string       `json:"title"`
	PollQuestion string       `json:"question"`
	PollOptions  []PollOption `json:"pollOptions"`
}
type PollList struct {
	Polls map[uint]Poll //A map of VoterIDs as keys and Voter structs as values
}
type cache struct {
	cacheClient *redis.Client
	jsonHelper  *rejson.Handler
	context     context.Context
}
type PollCache struct {
	cache
}

// constructor for VoterList struct
func NewPoll(id uint, title, question string) *Poll {
	return &Poll{
		PollID:       id,
		PollTitle:    title,
		PollQuestion: question,
		PollOptions:  []PollOption{},
	}
}

func NewpollOption(id uint, value string) *PollOption {
	return &PollOption{
		PollOptionID:    id,
		PollOptionValue: value,
	}
}
func NewSamplePoll() *Poll {
	return &Poll{
		PollID:       1,
		PollTitle:    "Favorite Pet",
		PollQuestion: "What type of pet do you like best?",
		PollOptions: []PollOption{
			{PollOptionID: 1, PollOptionValue: "Dog"},
			{PollOptionID: 2, PollOptionValue: "Cat"},
			{PollOptionID: 3, PollOptionValue: "Fish"},
			{PollOptionID: 4, PollOptionValue: "Bird"},
			{PollOptionID: 5, PollOptionValue: "NONE"},
		},
	}
}

func (p *Poll) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
func (p *PollOption) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
func NewPollCache() (*PollCache, error) {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}

	return NewWithCacheInstance(redisUrl)
}

// The constructor function that returns a pointer to a new VoterCache.
func NewWithCacheInstance(location string) (*PollCache, error) {
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

	return &PollCache{
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

func (vc *PollCache) getItemFromRedis(key string, item *Poll) error {
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
func (v *PollCache) AddPoll(poll Poll) error {
	if _, err := v.GetItem(poll.PollID); err == nil {
		return errors.New("poll already exists")
	}

	// v.polls[poll.pollID] = poll
	redisKey := redisKeyFromId(poll.PollID)
	if _, setErr := v.jsonHelper.JSONSet(redisKey, ".", poll); setErr != nil {
		return setErr
	}
	return nil
}
func (v *PollCache) AddpollOption(pollID, PollOptionID uint, pollOptionObject PollOption) (PollOption, error) {
	poll, err := v.GetItem(pollID)
	if err != nil {
		return PollOption{}, errors.New("poll does not exist")
	}

	for _, poll := range poll.PollOptions {
		if poll.PollOptionID == PollOptionID {
			return PollOption{}, errors.New("polloption already exist in this poll")
		}
	}

	newpollOption := pollOptionObject

	poll.PollOptions = append(poll.PollOptions, newpollOption)

	redisKey := redisKeyFromId(poll.PollID)

	if _, setErr := v.jsonHelper.JSONSet(redisKey, ".", poll); setErr != nil {
		return PollOption{}, setErr
	}

	return newpollOption, nil
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

func (v *PollCache) GetItem(id uint) (Poll, error) {
	// Check if the item exists before trying to get it
	// This is a good practice, return an error if the
	// item does not exist
	var poll Poll

	redisKey := redisKeyFromId(id)
	if err := v.getItemFromRedis(redisKey, &poll); err != nil {
		return Poll{}, errors.New("voter does not exist")
	}

	return poll, nil
}
func (v *PollCache) GetAllPolls() ([]Poll, error) {
	var polls []Poll

	pattern := fmt.Sprintf("%s*", RedisKeyPrefix)
	keys, err := v.cacheClient.Keys(v.context, pattern).Result()

	if err != nil {
		return polls, err
	}

	for _, key := range keys {
		var poll Poll

		err := v.getItemFromRedis(key, &poll)
		if err != nil {
			return polls, err
		}

		polls = append(polls, poll)
	}

	return polls, nil
}
func (v *PollCache) GetAllPollOptions(id uint) ([]PollOption, error) {
	poll, err := v.GetItem(id)
	if err != nil {
		return nil, errors.New("Voter not found")
	}
	return poll.PollOptions, nil
}

func (v *PollCache) GetPollOptionsByID(id uint, pollId uint) (PollOption, error) {
	voter, err := v.GetItem(id)
	if err != nil {
		return PollOption{}, errors.New("Voter not found")
	}

	for _, poll := range voter.PollOptions {
		if poll.PollOptionID == pollId {
			return poll, nil
		}
	}

	return PollOption{}, nil
}
