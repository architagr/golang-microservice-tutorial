package data

type Employee struct {
	ID         int    `json:"id" uri:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Role       []int  `json:"roles" form:"roles"`
	Department string `json:"department" form:"department"`
}
