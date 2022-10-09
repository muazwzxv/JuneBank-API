package request

type CreateAccount struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Month      int    `json:"month"`
	Day        int    `json:"day"`
	Year       int    `json:"year"`
}
