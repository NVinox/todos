package todos

type Todo struct {
	Id     int    `json="id"`
	Title  string `json="title"`
	Body   string `json="body"`
	UserId int    `json="userId"`
}
