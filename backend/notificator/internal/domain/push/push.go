package push

type Push struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Icon  string `json:"icon"`
}
