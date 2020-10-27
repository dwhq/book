package model

type List struct {
	Id uint `gorm:"<-:update"` // 允许读和更新
	Content string `gorm:"<-:update"` // 允许读和更新
	Update_at string `gorm:"<-:update"` // 允许读和更新
}
