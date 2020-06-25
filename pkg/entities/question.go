package entities

//Question depicts request structure of adding question
type Question struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
}
