package domain

type Users struct {
	ID        int    `json:"id"`
	Name      string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	UserType  string `json:"user_type"`
	State     string `json:"state"`
	Deteled   string `json:"deteled"`
	CreatedOn string `json:"created_on"`
}
