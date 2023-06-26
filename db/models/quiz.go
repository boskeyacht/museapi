package models

import (
	"context"
	"time"

	"github.com/boskeyacht/museapi/internal/types"
	"github.com/uptrace/bun"
)

type Quiz struct {
	ID        int64             `bun:"id,pk" json:"id"`
	AuthorID  int64             `bun:"author_id" json:"author_id"`
	CreatedAt time.Time         `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time         `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
	Title     string            `bun:"title,notnull" json:"title"`
	Questions []*types.Question `bun:"rel:has-many" json:"questions"`
	Score     int               `bun:"score,notnull" json:"score"`
}

func DefaultQuiz() *Quiz {
	return &Quiz{
		ID:        0,
		AuthorID:  0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Title:     "",
		Questions: []*types.Question{},
		Score:     0,
	}
}

func (q *Quiz) NewQuiz(ctx context.Context, db *bun.DB) error {
	res, err := db.NewInsert().Model(&Quiz{
		AuthorID:  q.AuthorID,
		Title:     q.Title,
		Questions: q.Questions,
		Score:     q.Score,
	}).Exec(ctx)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil
	}

	q.ID = id

	return nil
}

func (q *Quiz) GetQuiz(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(q).Where("id = ?", q.ID).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (q *Quiz) UpdateQuiz(ctx context.Context, db *bun.DB) error {
	_, err := db.NewUpdate().Model(q).Where("id = ?", q.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (q *Quiz) DeleteQuiz(ctx context.Context, db *bun.DB) error {
	_, err := db.NewDelete().Model(q).Where("id = ?", q.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
