package entities

//User represents user struct
type User struct {
	UUID            string
	EmailID         string
	SolvedQuestions []string
	HintsUsed       []string
}
