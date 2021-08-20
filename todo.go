package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	Userid int
	Listid int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:done`
}

type ListsItem struct {
	Id     int
	Listid int
	Itemid int
}
