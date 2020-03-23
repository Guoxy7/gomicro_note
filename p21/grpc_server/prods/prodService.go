package prods

import (
	"context"
	"gomicro_note/p21/grpc_server/models"
	"strconv"
	"time"
)

// ProdService 商品服务
type ProdService struct{}

func newProd(id int32, pname string) *models.ProdModel {
	return &models.ProdModel{ProdId: id, ProdName: pname}
}

// GetProdList 返回商品列表
func (*ProdService) GetProdList(ctx context.Context, in *models.ProdRequest, res *models.ProdListResponse) error {
	// 超时测试
	time.Sleep(time.Second * 3)
	models := make([]*models.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodName"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}

// GetProdDetail 获取单个商品
func (*ProdService) GetProdDetail(ctx context.Context, in *models.ProdRequest, res *models.ProdDetailResponse) error {
	// 超时测试
	time.Sleep(time.Second * 3)
	res.Data = newProd(in.ProdId, "测试商品")
	return nil
}
