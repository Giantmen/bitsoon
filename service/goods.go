package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Giantmen/bitsoon/common"
	"github.com/Giantmen/bitsoon/config"
	"github.com/Giantmen/bitsoon/log"
	"github.com/Giantmen/bitsoon/proto"
	"github.com/Giantmen/bitsoon/store"

	"github.com/gorilla/mux"
	"fmt"
)

type GoodsManager struct {
	db *store.Mysql
}

func NewGoodsManager(cfg *config.Config) (*GoodsManager, error) {
	log.Infof("%+v", *cfg)
	db, err := store.NewMysql(cfg.Mysql)
	if err != nil {
		return nil, err
	}
	return &GoodsManager{db: db}, nil
}

func (gm *GoodsManager) InsertHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("insert goods")

	gi := new(proto.GoodsInsert)
	err := common.ParseQuery(r, gi)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("insert into goods(location, price, restvolume, totalvolume, picture,uid)"+
		" values('%s',%v,%d,%d,'%s','%s')", gi.Location,gi.Price,gi.Restvolume,gi.Totalvolume,gi.Picture,gi.Uid)

	id,err := gm.db.Exec(sql, context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", id)
}

func (gm *GoodsManager) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("delete goods")

	gd := new(proto.GoodsDelete)
	err := common.ParseQuery(r, gd)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("delete from goods where id=%d",gd.Id)
	_,err = gm.db.Exec(sql,context.Background())
	if err !=nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	common.HttpResponse(w, 200, "success", nil)
}

func (gm *GoodsManager) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("update goods")

	gd := new(proto.GoodsUpdate)
	err := common.ParseQuery(r, gd)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	sql := fmt.Sprintf("update goods set location='%s',price=%v,restvolume=%d,totalvolume=%d,picture='%s',uid='%s' where id=%d",
	gd.Location,gd.Price,gd.Restvolume,gd.Totalvolume,gd.Picture,gd.Uid,gd.Id)
	_,err = gm.db.Exec(sql,context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", nil)

}

func (gm *GoodsManager) QueryAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("query all goods")

	sql := fmt.Sprintf("select * from goods")
	rrs := make([]*proto.Goods, 0)
	_, err := gm.db.Query(sql,&rrs,context.Background())
	if err != nil {
		common.HttpResponse(w, 500,  err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", rrs)
}

func (gm *GoodsManager) QueryOneHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("query one goods")

	vars := mux.Vars(r)
	goodsIDStr, ok := vars["goodsID"]
	if !ok {
		common.HttpResponse(w, 500, "goods id "+goodsIDStr+" not exist",nil)
		return
	}
	goodsID, err := strconv.ParseInt(goodsIDStr, 10, 64)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("select * from goods where id=%d",goodsID)
	rrs := make([]*proto.Goods, 0)
	num, err := gm.db.Query(sql,&rrs,context.Background())
	if err != nil {
		common.HttpResponse(w, 500,  err.Error(),nil)
		return
	}
	if num <=0 {
		common.HttpResponse(w, 404,  goodsIDStr+" not found",nil)
		return
	}
	common.HttpResponse(w, 200, "success", rrs[0])
}
