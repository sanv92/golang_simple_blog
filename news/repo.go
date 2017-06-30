package news

type News struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Alias       string `json:"alias" db:"alias"`
	Description string `json:"description" db:"description"`
	Content     string `json:"content" db:"content"`
}
