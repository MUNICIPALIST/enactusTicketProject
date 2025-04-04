package entity

type PrePost struct {
	ID      int    `json:"id"`
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Tickets int    `json:"tickets"`
}
