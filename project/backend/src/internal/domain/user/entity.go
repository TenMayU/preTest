package user

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
