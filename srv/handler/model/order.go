package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNo    string  `gorm:"type:varchar(32);"`
	UserID     int     `gorm:"not null;comment:用户ID"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null;comment:订单总金额"`
	PayStatus  int     `gorm:"default:0;comment:支付状态 0未支付 1已支付"`
}

func (o *Order) OrderAdd(db *gorm.DB) error {
	return db.Debug().Create(o).Error
}

func (o *Order) OrderItemAdd(db *gorm.DB, items []*OrderItem) error {
	return db.Debug().Create(items).Error
}

type OrderItem struct {
	gorm.Model
	OrderNo    string  `gorm:"index;not null;comment:订单编号"`
	GoodsID    uint    `gorm:"not null;comment:商品ID"`
	GoodsName  string  `gorm:"type:varchar(50);comment:商品名称"`
	GoodsPrice float64 `gorm:"type:decimal(10,2);comment:商品单价"`
	Num        int     `gorm:"not null;comment:购买数量"`
}

//type Order struct {
//	gorm.Model
//	OrderSn string `gorm:"type:varchar(30)"`
//	Name    string `gorm:"type:varchar(30)"`
//	Num     int    `gorm:"type:int(11)"`
//}

//func (o *Order) FindOrder(db *gorm.DB, name string) error {
//	return db.Debug().Where("name = ?", name).Find(&o).Error
//}
//
//func (o *Order) OrderAdds(db *gorm.DB) error {
//	return db.Debug().Create(&o).Error
//}
//func (o *Order) OrderList(db *gorm.DB, id int32) ([]*__.OrderList, []*__.OrderList) {
//	var list []*__.OrderList
//	db.Debug().Select("orders.*").Where("id = ?", id).Find(&list)
//	return list, nil
//
//}
