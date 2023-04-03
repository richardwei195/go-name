package models

import (
	"gorm.io/gorm"
)

// TWord 单词解析
type TWord struct {
	gorm.Model
	word         string
	pinyin       string
	bushou       string
	stokes       int32
	FiveElements FiveElementType
	weights      int32
	isUnCommon   bool
	meaning      string
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
