package authdto

type LoginResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
