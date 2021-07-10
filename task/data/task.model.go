package data

type Task struct {
	ID        int    `json:"id" uri:"id", form:"id"`
	Name      string `json:"name" form:"name"`
	Employee int  `json:"employee" form:"employee"`
}
