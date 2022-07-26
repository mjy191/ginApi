package Service

import (
	"encoding/json"
	"ginApi/Common/Enum"
	"ginApi/Common/Tools"
	"ginApi/Models"
)

var pageSize int64 = 5

type OrderService struct{}

type orderCopy struct {
	Models.Order
}

type OrderParam struct {
	Id     int   `form:"id" json:"id"`
	UserId int   `form:"userId" json:"userId"`
	Page   int64 `form:"page" json:"page" binding:"omitempty,min=1" min_msg:"page最小值1"`
	Status int   `form:"status" json:"status" binding:"omitempty,oneof=1 2 3" msg:"status值错误"`
}

type customerUser1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type customerOrder struct {
	Id        int            `json:"id"`
	UserId    int            `json:"userId"`
	Status    int            `json:"status"`
	StatusTxt string         `json:"statusTxt"`
	User      *customerUser1 `json:"user"`
	Cdt       Models.XTime   `json:"cdt"`
	Mdt       Models.XTime   `json:"mdt"`
}

func (this *orderCopy) MarshalJSON() ([]byte, error) {
	order := customerOrder{
		Id:        this.Id,
		UserId:    this.UserId,
		Status:    this.Status,
		Cdt:       this.Cdt,
		Mdt:       this.Mdt,
		StatusTxt: Tools.GetEnumValue(this.Status, Enum.Status),
	}
	if this.User != nil {
		order.User = &customerUser1{
			Id:   this.User.Id,
			Name: this.User.Name,
		}
	}
	return json.Marshal(order)
}

func (this OrderService) Lists(param *OrderParam) ([]*orderCopy, int64, int64, error) {
	var order Models.Order
	var result []*Models.Order
	query := Models.DB.Model(&order)

	//条件查询判断
	if param.Status > 0 {
		query = query.Where("status=?", param.Status)
	}
	//根据用户条件查询
	query = query.Where("userId=?", param.UserId)

	//计算总页
	var total int64
	var lastPage int64
	query.Count(&total)
	Tools.GetPage(total, &lastPage, &param.Page, pageSize)
	//获取数据
	query = query.Limit(int(pageSize)).Offset(int(pageSize * (param.Page - 1)))
	err := query.Preload("User").Find(&result).Error
	if err != nil {
		return nil, 0, 0, err
	}
	var list []*orderCopy
	//遍历数据把指针放进去
	for _, val := range result {
		list = append(list, &orderCopy{
			*val,
		})
	}
	return list, lastPage, total, nil
}
