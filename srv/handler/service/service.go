package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yueka/pkg"
	__ "yueka/proto"
	"yueka/srv/basic/config"
	"yueka/srv/handler/model"

	"github.com/google/uuid"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {

	var register model.User
	err := register.FindRegister(config.DB, in.Name)
	if err != nil {
		return &__.RegisterResp{
			Msg:  "不存在",
			Code: 400,
		}, nil
	}
	user := model.User{
		Name:     in.Name,
		Password: in.Password,
	}
	err = user.CreateAdd(config.DB)
	if err != nil {
		return &__.RegisterResp{
			Msg:  "注册失败",
			Code: 400,
		}, nil
	}
	return &__.RegisterResp{
		Msg:  "注册成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.LoginResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	if user.Password != in.Password {
		return &__.LoginResp{
			Msg:  "密码错误",
			Code: 400,
		}, nil
	}
	//handler, err := pkg.TokenHandler(strconv.Itoa(int(user.ID)))
	//if err != nil {
	//	return nil, err
	//}
	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) MiaoList(_ context.Context, in *__.MiaoListReq) (*__.MiaoListResp, error) {

	//var list []__.MiaoList
	var miao model.Miao
	err, _ := miao.FindMiao(config.DB, in.Id)
	if err != nil {
		return &__.MiaoListResp{
			Msg:  "列表查询失败",
			Code: 200,
		}, nil
	}

	return &__.MiaoListResp{
		Msg:  "列表展示成功",
		Code: 200,
		//Date: list,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {

	var goods model.Goods
	config.DB.Begin()
	err := goods.FinfGoods(config.DB, in.Name)
	if err != nil {
		config.DB.Rollback()
		return &__.GoodsAddResp{
			Msg:  "商品不存在",
			Code: 400,
		}, nil
	}
	m := model.Goods{
		Name:  in.Name,
		Price: float64(in.Price),
		Num:   int(in.Num),
	}
	err = m.GoodsAdd(config.DB)
	if err != nil {
		config.DB.Rollback()
		return &__.GoodsAddResp{
			Msg:  "商品添加失败",
			Code: 200,
		}, nil
	}
	config.DB.Commit()
	return &__.GoodsAddResp{
		Msg:  "商品添加成功",
		Code: 200,
	}, nil
}

//// SayHello implements helloworld.GreeterServer
//func (s *Server) OrderAdds(_ context.Context, in *__.OrderAddsReq) (*__.OrderAddsResp, error) {
//
//	var order model.Order
//	config.DB.Begin()
//	err := order.FindOrder(config.DB, in)
//	if err != nil {
//		config.DB.Rollback()
//		return &__.OrderAddsResp{
//			Msg:  "订单不存在",
//			Code: 400,
//		}, nil
//	}
//	m := model.Orders{
//
//	}
//	err = m.OrderAdd(config.DB)
//	if err != nil {
//		config.DB.Rollback()
//		return &__.OrderAddsResp{
//			Msg:  "订单添加失败",
//			Code: 400,
//		}, nil
//	}
//	//pkg.XueHua()
//	config.DB.Commit()
//	return &__.OrderAddsResp{
//		Msg:  "订单添加成功",
//		Code: 200,
//	}, nil
//}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderList(_ context.Context, in *__.OrderListReq) (*__.OrderListResp, error) {

	//var order model.Order
	////list, _ := order.OrderList(config.DB, in.Id)

	return &__.OrderListResp{
		Msg:  "订单列表查询成功",
		Code: 200,
		//Date: list,
	}, nil
}
func (s *Server) OrderAdd(_ context.Context, in *__.OrderAddReq) (*__.OrderAddResp, error) {
	timeStr := time.Now().Format("20060102150405")
	uuidStr := uuid.New().String()
	orderSn := fmt.Sprintf("%v%v", timeStr, uuidStr[:8])
	total := 0.0
	var orderItems []*model.OrderItem
	for _, item := range in.List {
		var goods model.Goods
		err := goods.FindGoodsById(config.DB, item.GoodsId)
		if err != nil {
			return nil, errors.New("商品不存在")
		}
		subTotal := goods.Price * float64(item.Quantity)
		total += float64(subTotal)
		orderItem := model.OrderItem{
			OrderNo:    orderSn,
			GoodsID:    goods.ID,
			GoodsName:  goods.Name,
			GoodsPrice: float64(goods.Price),
			Num:        int(item.Quantity),
		}
		orderItems = append(orderItems, &orderItem)

	}
	order := model.Order{
		OrderNo:    orderSn,
		UserID:     int(in.UserID),
		TotalPrice: total,
		PayStatus:  0,
	}
	err := order.OrderAdd(config.DB)
	if err != nil {
		return nil, errors.New("订单创建失败")
	}
	for i, _ := range orderItems {
		orderItems[i].ID = uint(int(order.ID))
	}
	err = order.OrderItemAdd(config.DB, orderItems)
	if err != nil {
		return nil, errors.New("明细添加失败")
	}
	pay := pkg.Alipay(orderSn, total)

	return &__.OrderAddResp{
		OrderSn: orderSn,
		Total:   float32(total),
		PayUrl:  pay,
	}, nil
}
