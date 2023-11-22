package middleware

type TokenParsed struct {
	Email string  `json:"email"`
	Exp   float64 `json:"exp"`
	Id    string  `json:"id"`
	Role  string  `json:"role"`
}
