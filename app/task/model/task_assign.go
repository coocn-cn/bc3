package model

import (
	"time"
)

// TaskAssign 任务 - 分配
type TaskAssign struct {
	// ID ID
	ID int32 `gorm:"column:id;primary_key" json:"id" `
	// TaskID 任务ID
	TaskID int32 `gorm:"column:task_id" json:"task_id" `
	// UserID 用户ID
	UserID int32 `gorm:"column:user_id" json:"user_id" `
	// OperatorID 操作人(创建人)ID
	OperatorID int32 `gorm:"column:operator_id" json:"operator_id" `
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" `
}

func (t TaskAssign) TableName() string {
	return "task_assign"
}
