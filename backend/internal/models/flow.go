type Flow struct {
	ID     uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Title  string    `json:"title"`
	Level  string    `json:"level"`
	Cover  string    `json:"cover"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User   User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
