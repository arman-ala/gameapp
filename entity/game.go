package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []uint // 2 Players in a game
	StartTime   time.Time // when the game starts
	// EndTime     time.Time // when the game ends
	// IsFinished  bool
}

// Player is the User but he/she is playing in a specific game
type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   int
	Answers []PlayerAnswer
}

// Player's answer to a Question instance
type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}
