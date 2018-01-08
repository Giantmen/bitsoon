package service

import (
	"context"
	"net/http"
	"strconv"
	"fmt"

	"github.com/Giantmen/bitsoon/common"
	"github.com/Giantmen/bitsoon/config"
	"github.com/Giantmen/bitsoon/log"
	"github.com/Giantmen/bitsoon/proto"
	"github.com/Giantmen/bitsoon/store"

	"github.com/gorilla/mux"
)

type UserManager struct {
	db *store.Mysql
}

func NewUserManager(cfg *config.Config) (*UserManager, error) {
	log.Infof("%+v", *cfg)
	db, err := store.NewMysql(cfg.Mysql)
	if err != nil {
		return nil, err
	}
	return &UserManager{db: db}, nil
}

func (um *UserManager) InsertHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("insert user")

	ui := new(proto.UserInsert)
	err := common.ParseQuery(r, ui)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("insert into user(name, password, phone, email)"+
		" values('%s','%s','%s','%s')", ui.Name,ui.Password,ui.Phone,ui.Email)

	id,err := um.db.Exec(sql, context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", id)
}

func (um *UserManager) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("delete user")

	ud := new(proto.UserDelete)
	err := common.ParseQuery(r, ud)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("delete from user where id=%d",ud.Id)
	_,err = um.db.Exec(sql,context.Background())
	if err !=nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	common.HttpResponse(w, 200, "success", nil)
}

func (um *UserManager) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("update user")

	uu := new(proto.UserUpdate)
	err := common.ParseQuery(r, uu)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	sql := fmt.Sprintf("update user set name='%s',password='%s',phone='%s',email='%s' where id=%d",
		uu.Name,uu.Password,uu.Phone,uu.Email,uu.Id)
	_,err = um.db.Exec(sql,context.Background())
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", nil)

}

func (um *UserManager) QueryAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("query all user")

	sql := fmt.Sprintf("select * from user")
	rrs := make([]*proto.User, 0)
	_, err := um.db.Query(sql,&rrs,context.Background())
	if err != nil {
		common.HttpResponse(w, 500,  err.Error(),nil)
		return
	}
	common.HttpResponse(w, 200, "success", rrs)
}

func (um *UserManager) QueryOneHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("query one user")

	vars := mux.Vars(r)
	userIDStr, ok := vars["userID"]
	if !ok {
		common.HttpResponse(w, 500, "user id "+userIDStr+" not exist",nil)
		return
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		common.HttpResponse(w, 500, err.Error(),nil)
		return
	}

	sql := fmt.Sprintf("select * from user where id=%d",userID)
	rrs := make([]*proto.User, 0)
	num, err := um.db.Query(sql,&rrs,context.Background())
	if err != nil {
		common.HttpResponse(w, 500,  err.Error(),nil)
		return
	}
	if num <=0 {
		common.HttpResponse(w, 404,  userIDStr+" not found",nil)
		return
	}
	common.HttpResponse(w, 200, "success", rrs[0])
}
