package cache

import (
	"encoding/json"
	"ffcs/api/utils"
	"ffcs/pkg/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Auth  *redis.Client
	Cache *redis.Client
}

func (r Redis) GetCache(key string) (string, error) {
	return r.Cache.Get(ctx, key).Result()
}
func (r Redis) SetCache(key string, value []byte) error {
	return r.Cache.Set(ctx, key, value, 5*time.Minute).Err()
}
func (r Redis) DeleteCache(key string) error {
	return r.Cache.Del(ctx, key).Err()
}

func (r Redis) GetAuthCache(key string) (models.Session, error) {
	var session models.Session
	data, err := r.Auth.Get(ctx, key).Result()
	if err != nil {
		return session, err
	}
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return session, err
	}
	return session, nil
}

func (r Redis) SetAuthCache(key string, value []byte) error {
	return r.Auth.Set(ctx, key, value, time.Hour).Err() // setting auth cache for 1 hour
}

func (r Redis) CreateSession(userid uuid.UUID, isAdmin bool, Branch string) (string, error) {
	key := utils.RandStr()
	c := models.Session{
		UserId:  userid,
		Valid:   true,
		IsAdmin: isAdmin,
		Branch:  Branch,
	}

	jsonData, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return key, r.SetAuthCache(key, jsonData)
}
