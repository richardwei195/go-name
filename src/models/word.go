package models

import (
	"gorm.io/gorm"
)

// TWord 单词解析
type TWord struct {
	id           uint `gorm:"primary"`
	Word         string
	Pinyin       string
	Bushou       string
	Stokes       int32
	FiveElements FiveElementType
	Weights      int32
	IsUnCommon   bool
	Meaning      string
}

// GetWords 获取单个字解释
func GetWords(wordQueryArray []string) (*[]TWord, error) {
	var word *[]TWord
	error := db.Where("word IN ?", wordQueryArray).Find(&word).Error

	if error != nil && error != gorm.ErrRecordNotFound {
		return nil, error
	}

	return word, error
}
