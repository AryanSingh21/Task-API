package connect

type Task struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Due_Date    string `gorm:"not null"`
	Status      string `gorm:"not null"`
}