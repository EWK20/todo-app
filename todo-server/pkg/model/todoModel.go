package model

type Todo struct {
	Id        int    `json:"id" gorm:"primaryKey" `
	Item      string `json:"item" `
	Completed bool   `json:"completed" `
}

func (u *Todo) TableName() string {
	// custom table name, this is default
	return "todo"
}
