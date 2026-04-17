package model

import (
	__ "yueka/proto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar(32)"`
}

func (u *User) FindRegister(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

func (u *User) CreateAdd(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

func (u *User) FindUser(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

type Miao struct {
	gorm.Model
	GoodsName string `gorm:"type:varchar(30)"`
	GoodsNum  int    `gorm:"type:int(11)"`
}

func (m Miao) FindMiao(db *gorm.DB, id int32) (error, []__.MiaoList) {
	var list []__.MiaoList
	db.Debug().
		Select("miaos.*").
		Where("id = ?", id).
		Find(&list)

	return nil, list
}

type Goods struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30)"`
	Price float64 `gorm:"type:decimal(10,2)"`
	Num   int     `gorm:"type:int(11)"`
}

func (g *Goods) FinfGoods(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}

func (g *Goods) GoodsAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (g *Goods) FindGoodsById(db *gorm.DB, id int64) interface{} {
	return db.Debug().Where("id = ?", id).Find(&g).Error
}
