package emulator

type Payment struct {
	Id				int			`json:"-"`
	UserId   		int 		`json:"userId"`
	UserEmail   	string 		`json:"userEmail"`
	Sum				float32 	`json:"sum"`
	Currency		string 		`json:"currency"`
	CreationDate	string 		`json:"-"` //date
	UpdateDate		string 		`json:"-"` //date
	Status			string		`json:"-"`
}
