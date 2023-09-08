package election

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

type Vote struct {
	VoteID    uint `json:"voteId"`
	VoterID   uint `json:"voterId"`
	PollID    uint `json:"pollId"`
	VoteValue uint `json:"voteValue"`
}
type VoteData struct {
	Votes []Vote //A map of VoterIDs as keys and Voter structs as values
}

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "vote:"
)

// constructor for VoterList struct
func NewVote(pid, vid, vtrid, vval uint) *Vote {
	return &Vote{
		VoteID:    vid,
		VoterID:   vtrid,
		PollID:    pid,
		VoteValue: vval,
	}
}

type cache struct {
	cacheClient *redis.Client
	jsonHelper  *rejson.Handler
	context     context.Context
}
type VoteCache struct {
	cache
}

func NewSampleVote() *Vote {
	return &Vote{
		VoteID:    1,
		PollID:    1,
		VoterID:   1,
		VoteValue: 1,
	}
}

func (p *Vote) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}
func NewVoteCache() (*VoteCache, error) {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}

	return NewWithCacheInstance(redisUrl)
}

// The constructor function that returns a pointer to a new VoterCache.
func NewWithCacheInstance(location string) (*VoteCache, error) {
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

	return &VoteCache{
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

func (vc *VoteCache) getItemFromRedis(key string, item *Vote) error {
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
func (v *VoteCache) AddVote(vote Vote) error {
	if _, err := v.GetItem(vote.VoteID); err == nil {
		return errors.New("vote already exists")
	}

	// v.polls[poll.pollID] = poll
	redisKey := redisKeyFromId(vote.PollID)
	if _, setErr := v.jsonHelper.JSONSet(redisKey, ".", vote); setErr != nil {
		return setErr
	}
	return nil
}
func (v *VoteCache) GetItem(id uint) (Vote, error) {
	// Check if the item exists before trying to get it
	// This is a good practice, return an error if the
	// item does not exist
	var Vote Vote

	redisKey := redisKeyFromId(id)
	if err := v.getItemFromRedis(redisKey, &Vote); err != nil {
		return Vote, errors.New("voter does not exist")
	}

	return Vote, nil
}
func (v *VoteCache) GetAllVotes() ([]Vote, error) {
	var Votes []Vote

	pattern := fmt.Sprintf("%s*", RedisKeyPrefix)
	keys, err := v.cacheClient.Keys(v.context, pattern).Result()

	if err != nil {
		return Votes, err
	}

	for _, key := range keys {
		var vote Vote

		err := v.getItemFromRedis(key, &vote)
		if err != nil {
			return Votes, err
		}

		Votes = append(Votes, vote)
	}

	return Votes, nil
}
