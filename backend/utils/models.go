package utils

import (
	"time"
        "sql"
)

type SessionData struct {
	ID       int
	Username string
	Role     string
	Expiry   time.Time
}

type Thread struct {
	ID         int    `sql:"id"`
	ThreadName string `sql:"thread_name"`
}

type Voting struct {
	ID       int    `sql:"id"`
	ThreadId int    `sql:"thread_id"`
	Title    string `sql:"title"`
	Descr    string `sql:"descr"`
}

type Vote struct {
    ID       int    `sql:"id"`
    UserId   int    `sql:"user_id"`
    VotingId int    `sql:"voting_id"`
    Vote     string `sql:"vote"`
}

type Option struct {
    VotingId   int    `sql:"voting_id"`
    OptionText string `sql:"option_text"`
    VoteCount  int    `sql:"vote_count"`
}

type Comment struct {
    ID          int    `sql:"id"`
    UserId      int    `sql:"user_id"`
    VotingId    int    `sql:"voting_id"`
    CommentText string `sql:"comment_text"`
}

type ContextKey string

