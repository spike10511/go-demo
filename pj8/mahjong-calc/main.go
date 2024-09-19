package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

type MahjongRequest struct {
	PlayerID    string `json:"player_id"`
	HuType      string `json:"hu_type"`   // "ron" or "tsumo"
	TargetID    string `json:"target_id"` // Only needed for "ron"
	LiBaoCount  int    `json:"li_bao_count"`
	ChiBaoCount int    `json:"chi_bao_count"`
	IsYakuman   bool   `json:"is_yakuman"`
}

func initializePlayerScore(playerID string) {
	playerKey := "player:" + playerID
	if rdb.HExists(ctx, playerKey, "score").Val() == false {
		rdb.HSet(ctx, playerKey, "score", 20) // Initialize with 20 points
	}
}

func calculate(c *gin.Context) {
	var req MahjongRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	playerKey := "player:" + req.PlayerID
	targetKey := "player:" + req.TargetID

	// Initialize scores if not already
	initializePlayerScore(req.PlayerID)
	if req.HuType == "ron" {
		initializePlayerScore(req.TargetID)
	}

	// Calculate points based on HuType and Yakuman
	if req.IsYakuman {
		if req.HuType == "ron" {
			rdb.HIncrBy(ctx, playerKey, "score", 10)
			rdb.HIncrBy(ctx, targetKey, "score", -10)
		} else if req.HuType == "tsumo" {
			rdb.HIncrBy(ctx, playerKey, "score", 30)
			for i := 1; i <= 4; i++ {
				otherKey := "player:" + strconv.Itoa(i)
				if otherKey != playerKey {
					rdb.HIncrBy(ctx, otherKey, "score", -10)
				}
			}
		}
	} else {
		if req.HuType == "ron" {
			rdb.HIncrBy(ctx, playerKey, "score", int64(req.LiBaoCount))
			rdb.HIncrBy(ctx, targetKey, "score", -int64(req.LiBaoCount))
		} else if req.HuType == "tsumo" {
			rdb.HIncrBy(ctx, playerKey, "score", int64(req.LiBaoCount)*3)
			for i := 1; i <= 4; i++ {
				otherKey := "player:" + strconv.Itoa(i)
				if otherKey != playerKey {
					rdb.HIncrBy(ctx, otherKey, "score", -int64(req.LiBaoCount))
				}
			}
		}
	}

	// Add ChiBao points
	if req.ChiBaoCount > 0 {
		rdb.HIncrBy(ctx, playerKey, "score", int64(req.ChiBaoCount))
	}

	// Get all player scores
	playerScores := make(map[string]string)
	for i := 1; i <= 4; i++ {
		playerKey := "player:" + strconv.Itoa(i)
		score, _ := rdb.HGet(ctx, playerKey, "score").Result()
		playerScores[playerKey] = score
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Scores updated",
		"scores":  playerScores,
	})
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := gin.Default()
	router.POST("/calculate", calculate)
	router.Run(":8081")
}
