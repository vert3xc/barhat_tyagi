package utils

import (
	"time"
)

type SessionData struct {
	ID       int
	Username string
	Role     string
	Expiry   time.Time
}

type Option struct {
	VotingId int
	OptionText string
	VoteCount int
}

type FullVoting struct{
	Voting Voting
	Options []Option
}

type Thread struct {
	ID         int
	ThreadName string
}

type Voting struct {
	ID       int
	ThreadId int
	Title    string
	Descr    string
}

type ContextKey string
