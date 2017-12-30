package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/linchongky/btcsoon/common"
	"github.com/linchongky/btcsoon/config"
	"github.com/linchongky/btcsoon/log"
	"github.com/linchongky/btcsoon/proto"
	"github.com/linchongky/btcsoon/store"

	"github.com/gorilla/mux"
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

	id, err := gm.db.InsertGoods(gi, context.Background())
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

	err = gm.db.DeleteGoods(gd.ID, context.Background())
	if err != nil {
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

	err = gm.db.UpdateGoods(gd, context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", nil)

}

func (gm *GoodsManager) QueryAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("query all goods")
	//gqa := new(proto.GoodsQueryAll)
	//err = common.ParseQuery(r, gqa)
	//if err != nil {
	//	return
	//}

	allg, err := gm.db.QueryAllGoods(context.Background())
	if err != nil {
		common.HttpResponse(w, 500,  err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", allg)
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
	//gqo := new(proto.GoodsQueryOne)
	//err = common.ParseQuery(r, gqo)
	//if err != nil {
	//	return
	//}

	g, err := gm.db.QueryOneGoods(goodsID, context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", g)
}
