package todos

type Todo struct {
	ID          int    `json:"id"`
	Created_at  string `json:"createdAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}