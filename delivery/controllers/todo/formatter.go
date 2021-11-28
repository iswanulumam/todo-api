package todo

type todoRequest struct {
	Title string `json:"title" form:"title"`
}
