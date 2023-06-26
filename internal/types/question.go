package types

type Question string

const (
	TrueFalseQuestion      Question = "TrueFalseQuestion"
	MultipleChoiceQuestion Question = "MultipleChoiceQuestion"
	ShortAnswerQuestion    Question = "ShortAnswerQuestion"
)
