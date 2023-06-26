package models

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`
	ID            int64        `bun:"id,pk" json:"id"`
	Username      string       `bun:"username,notnull,unique" json:"username"`
	CreatedAt     time.Time    `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time    `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
	FirstName     string       `bun:"first_name,notnull" json:"first_name"`
	LastName      string       `bun:"last_name,notnull" json:"last_name"`
	Email         string       `bun:"email,notnull" json:"email"`
	Password      string       `bun:"password,notnull" json:"password"`
	Quizzes       []*Quiz      `bun:"rel:has-many" json:"quizzes"`
	Flashcards    []*Flashcard `bun:"rel:has-many" json:"flashcards"`
}

func DefaultUser() *User {
	return &User{
		ID:         0,
		Username:   "",
		FirstName:  "",
		LastName:   "",
		Email:      "",
		Password:   "",
		Quizzes:    []*Quiz{},
		Flashcards: []*Flashcard{},
	}
}

// @todo relatioonships
func (u *User) NewUser(ctx context.Context, db *bun.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(hashed)

	res, err := db.NewInsert().Model(&User{
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  encoded,
	}).Exec(ctx)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

func (u *User) GetUserByUsername(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(u).Where("username = ?", u.Username).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByID(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(u).WherePK().Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) CheckPassword(password string) bool {
	hashed, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))

	return err == nil
}

func (u *User) UpdateUser(ctx context.Context, db *bun.DB) error {
	_, err := db.NewUpdate().Model(u).Where("id = ?", u.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) AppendQuiz(ctx context.Context, db *bun.DB, quiz *Quiz) error {
	err := db.NewSelect().Model(u).Where("username = ?", u.Username).Scan(ctx)
	if err != nil {
		return err
	}

	_, err = db.NewUpdate().Model(u).Set("quizzes = ?", u.Quizzes).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) AppendFlashcards(ctx context.Context, db *bun.DB, flashcards *Flashcard) error {
	err := db.NewSelect().Model(u).Where("username = ?", u.Username).Scan(ctx)
	if err != nil {
		return err
	}

	_, err = db.NewUpdate().Model(u).Set("flashcards = ?", u.Flashcards).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, db *bun.DB) error {
	_, err := db.NewDelete().Model(u).Where("id = ?", u.ID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetQuizzes(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(&u.Quizzes).Where("user_id = ?", u.ID).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetFlashcards(ctx context.Context, db *bun.DB) error {
	err := db.NewSelect().Model(&u.Flashcards).Where("user_id = ?", u.ID).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}
