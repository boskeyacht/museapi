package types

type Question interface {
	GetQuestion() string
	GetCorrectAnswer() string
	GetPossibleAnswers() []string
	GetQuestionType() string
}
