package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Order struct {
	Id     int   `json:"id" gorm:"column:id;primaryKey"`
	UserId int   `json:"userId" gorm:"column:userId"`
	Status int   `json:"status" gorm:"column:status"`
	User   *User `json:"user" gorm:"foreignKey:userId;association_foreignKey:id"`
	Cdt    XTime `json:"cdt" gorm:"column:cdt;default:null"`
	Mdt    XTime `json:"mdt" gorm:"column:mdt;default:null"`
}

func (Order) TableName() string {
	return "order"
}

type XTime struct {
	time.Time
}

// 3.为Xtime实现Value方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4.为Xtime实现Scan方法，读取数据库时会调用该方法将时间数据转换成自定义的时间类型
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not conver %v to timestamp", v)
}

func (t XTime) MarshalJSON() ([]byte, error) {
	if &t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}
