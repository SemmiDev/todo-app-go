package model

// TodoData is a struct for todo data
type TodoData struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status bool   `json:"status" db:"status"`
}

// Clone returns a copy of the todo data
func (data *TodoData) Clone() *TodoData {
	return &TodoData{
		ID:     data.ID,
		Title:  data.Title,
		Status: data.Status,
	}
}
