package models

import "time"

type Flashcard struct {
	ID        int64     `bun:"id,pk"`
	AuthorID  int64     `bun:"author_id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Question  string    `bun:"question,notnull"`
	Answer    string    `bun:"answer,notnull"`
}
