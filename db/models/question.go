package models

import (
	"context"
	"time"

	"github.com/boskeyacht/museapi/internal/types"
	"github.com/uptrace/bun"
)

type Question struct {
	ID              int64          `bun:"id,pk" json:"id"`
	CreatedAt       time.Time      `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time      `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
	Type            types.Question `bun:"type,notnull" json:"type"`
	QuizID          int64          `bun:"quiz_id" json:"quiz_id"`
	Question        string         `bun:"question,notnull" json:"question"`
	Answer          bool           `bun:"answer,notnull" json:"answer"`
	PossibleAnswers []string       `bun:"possible_answers,notnull" json:"possible_answers"`
	IsCorrect       bool           `bun:"is_correct,notnull" json:"is_correct"`
}

func DefaultQuestion() *Question {
	return &Question{
		ID:        0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "",
		QuizID:    0,
		Question:  "",
		Answer:    false,
		IsCorrect: false,
	}
}

func (q *Question) NewQuestion(ctx context.Context, db *bun.DB) error {
	res, err := db.NewInsert().Model(&Question{
		Type:            q.Type,
		QuizID:          q.QuizID,
		Question:        q.Question,
		Answer:          q.Answer,
		PossibleAnswers: q.PossibleAnswers,
		IsCorrect:       q.IsCorrect,
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

func (q *Question) GetQuestion(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(q).Where("id = ?", q.ID).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (q *Question) UpdateQuestion(ctx context.Context, db *bun.DB) error {
	res, err := db.NewUpdate().Model(&Question{
		ID:              q.ID,
		Type:            q.Type,
		QuizID:          q.QuizID,
		Question:        q.Question,
		Answer:          q.Answer,
		PossibleAnswers: q.PossibleAnswers,
		IsCorrect:       q.IsCorrect,
	}).Where("id = ?", q.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (q *Question) DeleteQuestion(ctx context.Context, db *bun.DB) error {
	res, err := db.NewDelete().Model(&Question{
		ID: q.ID,
	}).Where("id = ?", q.ID).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
