package models

import (
	"gorm.io/gorm"
)

type SexType string

type FiveElementType string

const (
	Male     SexType         = "male"
	Female   SexType         = "female"
	TuShui   FiveElementType = "土水"
	ShuiTu   FiveElementType = "水土"
	TuMu     FiveElementType = "土木"
	MuTu     FiveElementType = "木土"
	TuHuo    FiveElementType = "土火"
	HuoTu    FiveElementType = "火土"
	MuShui   FiveElementType = "木水"
	ShuiMu   FiveElementType = "水木"
	MuHuo    FiveElementType = "木火"
	HuoMu    FiveElementType = "火木"
	JinShui  FiveElementType = "金水"
	ShuiJin  FiveElementType = "水金"
	JinMu    FiveElementType = "金木"
	MuJin    FiveElementType = "木金"
	JinTu    FiveElementType = "金土"
	TuJin    FiveElementType = "土金"
	JinHuo   FiveElementType = "金火"
	HuoJin   FiveElementType = "火金"
	ShuiHuo  FiveElementType = "水火"
	HuoShui  FiveElementType = "火水"
	Tu       FiveElementType = "土"
	TuTu     FiveElementType = "土土"
	Mu       FiveElementType = "木"
	MuMu     FiveElementType = "木木"
	Jin      FiveElementType = "金"
	JinJin   FiveElementType = "金金"
	Shui     FiveElementType = "水"
	ShuiShui FiveElementType = "水水"
	Huo      FiveElementType = "火"
	HuoHuo   FiveElementType = "火火"
)

type TName struct {
	//gorm.Model
	ID           uint            `gorm:"primary"`
	Name         string          `json:"name"`
	Sex          SexType         `json:"sex"`
	MachineScore float64         `json:"machineScore"`
	ManualScore  float64         `json:"manualScore"`
	FiveElements FiveElementType `json:"fiveElements"`
}

func GetNames(pageNum, pageSize int, maps interface{}) ([]*TName, error) {
	var names []*TName
	ret := db.Where(maps).Offset(pageNum).Limit(pageSize).Order("RAND()").Find(&names)

	if ret.Error != nil && ret.Error != gorm.ErrRecordNotFound {
		return nil, ret.Error
	}

	return names, ret.Error
}
