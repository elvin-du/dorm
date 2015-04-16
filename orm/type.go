package orm

type Users struct {
	Id     string `mysql:"id" json:"id"`
	Name   string `mysql:"name" json:"name"`
	Age    byte   `mysql:"age" json:"age"`
	Status byte   `mysql:"status"`
}
