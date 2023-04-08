package models

import (
	"gorm.io/gorm"
)

// TWord 单词解析
type TWord struct {
	id           uint            `gorm:"primary" json:"id"`
	Word         string          `json:"word"`
	Pinyin       string          `json:"pinyin"`
	Bushou       string          `json:"bushou"`
	Stokes       int32           `json:"stokes"`
	FiveElements FiveElementType `json:"fiveElements"`
	Weights      int32           `json:"weights"`
	IsUnCommon   bool            `json:"isUnCommon"`
	Meaning      string          `json:"meaning"`
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
