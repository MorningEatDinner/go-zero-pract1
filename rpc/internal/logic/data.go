package logic

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "木今",
		Phone: "13800001111",
	},
	"2": {
		Id:    "2",
		Name:  "小森",
		Phone: "15688880000",
	},
}
