package todo

type TodoRequestFormat struct {
	Title string `json:"title" form:"title"`
}
