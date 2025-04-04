// internal/entity/post.go
package entity

type Post struct {
	ID      int    `json:"id"`
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Tickets int    `json:"tickets"`
	Data    int64  `json:"data"`
	Status  string `json:"status"`
}
