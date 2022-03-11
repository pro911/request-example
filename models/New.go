package models

type New struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Author  string `json:"author" gorm:"index"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

func GetNews(pageNum int, pageSize int, maps interface{}) (news []New) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&news)
	return
}

func GetNewTotal(maps interface{}) (count int64) {
	db.Model(&New{}).Where(maps).Count(&count)
	return
}

func ExistNewByTitle(title string) bool {
	var new New
	db.Select("id").Where("title = ?", title).First(&new)
	if new.ID > 0 {
		return true
	}
	return false
}

func AddNew(title string, author string, url string, content string) bool {
	db.Create(&New{
		Title:   title,
		Author:  author,
		Url:     url,
		Content: content,
	})
	return true
}

func GetNew(id int) (new New) {
	db.Where("id = ?", id).First(&new)
	return
}
