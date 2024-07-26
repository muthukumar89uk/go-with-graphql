package dbmodel

type User struct {
	UserID      int     `json:"userid" gorm:"primaryKey"`
	Username    string  `json:"username" gorm:"not null"`
	Email       string  `json:"email" gorm:"not null"`
	Password    string  `json:"password" gorm:"not null"`
	PhoneNumber string  `json:"phone_number"`
	Profile     Profile `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Profile struct {
	ProfileID int    `json:"profileid" gorm:"column:profileid;primaryKey"`
	UserID    uint   `json:"userid" gorm:"not null"`
	FullName  string `json:"full_name" gorm:"not null"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
}

type Career struct {
	CareerID    uint   `json:"career_id" gorm:"primaryKey;autoIncrement"`
	Company     string `json:"company" gorm:"not null"`
	Position    string `json:"position" gorm:"not null"`
	JobType     string `json:"job_type"`
	Description string `json:"description"`
}

