package article

import (
	"github.com/zxzhaoqiqi/goblog/model"
	"github.com/zxzhaoqiqi/goblog/pkg/logger"
	"github.com/zxzhaoqiqi/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToUint64(idstr)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil

}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)

	if err = result.Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)

	if err = result.Error; err != nil {
		logger.Error(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}
