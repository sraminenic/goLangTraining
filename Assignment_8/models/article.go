package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID int
	Title string
	Content string
}

//create a Article
func CreateArticle(db *gorm.DB, Article *Article) (err error) {
	err = db.Create(Article).Error
	if err != nil {
		return err
	}
	return nil
}

//get all Article
func GetArticles(db *gorm.DB, Article *[]Article) (err error) {
	err = db.Find(Article).Error
	if err != nil {
		return err
	}
	return nil
}

//get Article by id
func GetArticle(db *gorm.DB, Article *Article, id string) (err error) {
	err = db.Where("id = ?", id).First(Article).Error
	if err != nil {
		return err
	}
	return nil
}

//update Article
func UpdateArticle(db *gorm.DB, Article *Article) (err error) {
	db.Save(Article)
	return nil
}

//delete Article
func DeleteArticle(db *gorm.DB, Article *Article, id string) (err error) {
	db.Where("id = ?", id).Delete(Article)
	return nil
}



