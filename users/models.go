package users

type Users struct {
	Id int `form: "id" json: "id"`
	Name string `form: "name" json: "name"`
	//Username string `form: "username" json: "username"`
}
