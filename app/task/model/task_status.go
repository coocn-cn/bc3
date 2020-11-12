package model

import (
	"time"
)

// TaskStatus 任务 - 状态
type TaskStatus struct {
	// ID ID
	ID int32 `gorm:"column:id;primary_key" json:"id" `
	// Name 名称
	Name string `gorm:"column:name" json:"name" `
	// Type 类型 - 通常，结束
	Type int32 `gorm:"column:type" json:"type" `
	// OperatorID 操作人(创建人)ID
	OperatorID int32 `gorm:"column:operator_id" json:"operator_id" `
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" `
}

func (t TaskStatus) TableName() string {
	return "task_status"
}
