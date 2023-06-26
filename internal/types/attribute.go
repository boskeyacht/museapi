package types

type Attribute struct {
	Original    string `json:"original"`
	Replacement string `json:"replacement"`
}

func NewAttribute(original, replacement string) *Attribute {
	return &Attribute{
		Original:    original,
		Replacement: replacement,
	}
}
