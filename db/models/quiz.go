package models

import (
	"time"

	"github.com/boskeyacht/museapi/internal/types"
)

type Quiz struct {
	ID        int64             `bun:"id,pk"`
	AuthorID  int64             `bun:"author_id"`
	CreatedAt time.Time         `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time         `bun:",nullzero,notnull,default:current_timestamp"`
	Title     string            `bun:"title,notnull"`
	Questions []*types.Question `bun:"rel:has-many"`
}
