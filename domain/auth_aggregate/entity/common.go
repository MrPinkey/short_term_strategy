package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Res struct {
	Data struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
		Token string `json:"token"`
	} `json:"data"`
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
}

type Menu struct {
	Id       int    `json:"id"`
	AuthName string `json:"authName"`
	Path     string `json:"path"`
	Children []Menu `json:"children"`
	Order    int    `json:"order"`
	Meta     struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
}
type Result struct {
	Data []Menu `json:"data"`
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
}
