package collections

type MovieComment struct {
	tableName struct{} `pg:"movie_comments"`
	Id        int      `json:"id" pg:"id,pk"`
	MovieId   string   `json:"movie_id" pg:"movie_id"`
	Email     string   `json:"email" pg:"email"`
	Comment   string   `json:"comment" pg:"comment"`
}
