package model

// Status use for user status, admin or user
type Status struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Status string `gorm:"unique" json:"status"`
}
