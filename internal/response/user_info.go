package response

type UserInfoShowRes struct {
	Id        string `json:"id"`
	Passport  string `json:"passport"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
