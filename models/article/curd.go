package article

import (
	"github.com/zxzhaoqiqi/goblog/models"
	"github.com/zxzhaoqiqi/goblog/pkg/logger"
	"github.com/zxzhaoqiqi/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToUint64(idstr)

	if err := models.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil

}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := models.DB.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (article *Article) Create() (err error) {
	result := models.DB.Create(&article)

	if err = result.Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := models.DB.Save(&article)

	if err = result.Error; err != nil {
		logger.Error(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := models.DB.Delete(&article)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}
