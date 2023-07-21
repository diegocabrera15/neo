package types

type UserTable struct {
	UserId         string `json:"userId"`
	FirstName      string `json:"firstName"`
	SurName        string `json:"surName"`
	Age            int    `json:"age"`
	Identification string `json:"identification"`
}
