package model

import (
	"time"
)

// Task 任务
type Task struct {
	// ID ID
	ID int32 `gorm:"column:id;primary_key" json:"id" `
	// ParentID 父任务ID
	ParentID int32 `gorm:"column:parent_id" json:"parent_id" `
	// Title 标题
	Title string `gorm:"column:title" json:"title" `
	// Content 内容
	Content string `gorm:"column:content" json:"content" `
	// StartTime 起始日期(可选)
	StartTime string `gorm:"column:start_time" json:"start_time" `
	// EndTime 截止日期(可选)
	EndTime string `gorm:"column:end_time" json:"end_time" `
	// ManHour 预估工时(可选)
	ManHour int32 `gorm:"column:man_hour" json:"man_hour" `
	// StatusID 状态ID
	StatusID int32 `gorm:"column:status_id" json:"status_id" `
	// OperatorID 操作人(创建人)ID
	OperatorID int32 `gorm:"column:operator_id" json:"operator_id" `
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" `
}

func (t Task) TableName() string {
	return "task"
}
