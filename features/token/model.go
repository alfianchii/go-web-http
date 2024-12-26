package token

import (
	"encoding/json"
	"fmt"
	"time"
	"web-http/config"
)

type Token struct {
	Username string `bson:"username" json:"username"`
	Value string `bson:"value,omitempty" json:"value"`
	IsBlacklist bool `bson:"isBlacklist" json:"isBlacklist"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

var (
	RedisSessionName = "session:"
	RedisSessionBlacklist = "blacklist:"
)

func GetValidTokenFromUser(username string) (Token, error) {
	var _, cancel = config.CtxTime()
	defer cancel()

	var token Token
	sessionKey := RedisSessionName + username
	tokenJSON, err := config.RedisClient.Get(config.CtxBg(), sessionKey).Result()
	if err != nil {
		return token, fmt.Errorf("unauthorized; token was blacklisted")
	}
	_ = json.Unmarshal([]byte(tokenJSON), &token)

	return token, nil
}

func CreateToken(username string, jwtToken string) (Token, error) {
	var _, cancel = config.CtxTime()
	defer cancel()

	var token Token = Token{
		Username: username,
		Value: jwtToken,
		IsBlacklist: false,
		CreatedAt: time.Now(),
	}

	sessionKey := RedisSessionName + username
	storedTokenJSON, _ := json.Marshal(token)
	err := config.RedisClient.Set(config.CtxBg(), sessionKey, storedTokenJSON, config.TokenDuration).Err()
	if err != nil {
		return token, fmt.Errorf("unauthorized; error storing token")
	}

	return token, nil
}

func BlacklistUsedToken(username string) error {
	var _, cancel = config.CtxTime()
	defer cancel()
	
	var sessionData Token
	sessionKey := RedisSessionName + username
	sessionDataJSON, _ := config.RedisClient.Get(config.CtxBg(), sessionKey).Result()
	_ = json.Unmarshal([]byte(sessionDataJSON), &sessionData)
	sessionData.IsBlacklist = true
	
	if sessionData.Value != "" {
		config.RedisClient.Del(config.CtxBg(), sessionKey)
		
		storedDataJSON, _ := json.Marshal(sessionData)
		sessionKey = RedisSessionBlacklist + sessionData.Value
		err := config.RedisClient.Set(config.CtxBg(), sessionKey, storedDataJSON, config.TokenDuration).Err()
		if err != nil {
			return fmt.Errorf("unauthorized; error blacklisting token")
		}
	}

	return nil
}