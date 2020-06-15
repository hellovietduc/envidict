package internal

type Word struct {
	Spelling      string        `json:"word"`
	Pronunciation string        `json:"pronunciation"`
	Definitions   []*Definition `json:"definitions"`
}

type Definition struct {
	Kind         string         `json:"kind,omitempty"`
	Descriptions []*Description `json:"descriptions"`
}

type Description struct {
	Meaning string `json:"meaning,omitempty"`
	Example string `json:"example,omitempty"`
}
