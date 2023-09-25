package database

type DBUserResponse struct {
	User struct {
		Props map[string]interface{} `json:"Props"`
	} `json:"user"`
}
