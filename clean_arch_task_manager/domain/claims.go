package domain

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}