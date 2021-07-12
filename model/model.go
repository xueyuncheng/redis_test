package model

import (
	"fmt"
	"redis_test/internal/log"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt JSONTime       `json:"created_at"`
	UpdatedAt JSONTime       `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type JSONTime time.Time

func (j *JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%v\"", time.Time(*j).Format("2006-01-02 15:04:06"))
	return []byte(stamp), nil
}

func (j *JSONTime) UnmarshalJSON(data []byte) error {
	tmp, err := time.ParseInLocation("\"2006-01-02 15:04:05\"", string(data), time.Local)
	if err != nil {
		log.Sugar.Error("时间解析错误", zap.Error(err))
		return fmt.Errorf("时间解析错误 %w", err)
	}

	*j = JSONTime(tmp)
	return nil
}
