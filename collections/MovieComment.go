package collections

type MovieComment struct {
	tableName struct{} `pg:"movie_comments"`
	ID        string   `json:"id" pg:"id,pk"`
	MovieId   int      `json:"movie_id" pg:"movie_id"`
	Email     string   `json:"email" pg:"email"`
	Comment   string   `json:"comment" pg:"comment"`
}
