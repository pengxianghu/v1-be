package defs

type Schedule struct {
	Id        int    `json:"id"`
	UserId    string `json:"user_id"`
	Topic     string `json:"topic"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Status    int    `json:"status"`
	Active    int    `json:"active"`
}
