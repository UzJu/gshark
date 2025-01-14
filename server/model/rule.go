// 自动生成模板Rule
package model

import (
	"github.com/madneal/gshark/global"
)

type Rule struct {
	global.GVA_MODEL
	RuleType string `json:"ruleType" form:"ruleType" gorm:"costringype;comment:;rule_type:varchar(200);size:200;"`
	Content  string `json:"content" form:"content" gorm:"column:content;comment:;type:varchar(100);size:100;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(100);size:100;"`
	Desc     string `json:"desc" form:"desc" gorm:"column:desc;comment:;type:varchar(300);size:300;"`
	Status   bool   `json:"status" form:"status" gorm:"column:status;comment:;type:int;size:3;"`
}

func (Rule) TableName() string {
	return "rule"
}
