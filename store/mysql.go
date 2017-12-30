package store

import (
	"context"
	"database/sql"
	"github.com/Giantmen/bitsoon/log"

	"github.com/Giantmen/bitsoon/config"
	"github.com/Giantmen/bitsoon/proto"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db  *sql.DB
	Cfg *config.Mysql
}

func NewMysql(cfg *config.Mysql) (*Mysql, error) {
	db, err := sql.Open("mysql", cfg.ConnStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.Ping()
	return &Mysql{
		db:  db,
		Cfg: cfg,
	}, nil
}

func (db *Mysql) QueryAllGoods(ctx context.Context) (*proto.GoodsQueryAllResult, error) {
	rows, err := db.db.Query("SELECT * FROM goods")
	if err != nil {
		return nil, err
	}
	rrs := new(proto.GoodsQueryAllResult)
	rrs.Gs = make([]*proto.Goods, 0)
	for rows.Next() {
		rr := new(proto.Goods)
		err = rows.Scan(&rr.ID, &rr.Location, &rr.Price, &rr.RestVolume, &rr.TotalVolume, &rr.Picture, &rr.UID)
		if err != nil {
			return nil, err
		}
		rrs.Gs = append(rrs.Gs, rr)
	}
	return rrs, nil
}

func (db *Mysql) QueryOneGoods(GoodsID int64, ctx context.Context) (*proto.GoodsQueryOneResult, error) {
	stmt, err := db.db.Prepare("SELECT * FROM goods WHERE id=?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(GoodsID)
	rr := new(proto.GoodsQueryOneResult)
	err = row.Scan(&rr.ID, &rr.Location, &rr.Price, &rr.RestVolume, &rr.TotalVolume, &rr.Picture, &rr.UID)
	if err != nil {
		return nil, err
	}
	log.Debugf("record %d : %+v", GoodsID, *rr)
	return rr, nil
}

func (db *Mysql) InsertGoods(gir *proto.GoodsInsert, ctx context.Context) (int64, error) {
	stmt, err := db.db.Prepare("INSERT goods SET location=?,price=?,restvolume=?,totalvolume=?,picture=?,uid=?")
	if err != nil {
		return 0, err
	}

	//createTime := time.Now().Format("2006-01-02 15:04:05")
	//updateTime := createTime
	res, err := stmt.Exec(gir.Location, gir.Price, gir.RestVolume, gir.TotalVolume, gir.Picture, gir.UID)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

func (db *Mysql) UpdateGoods(gu *proto.GoodsUpdate, ctx context.Context) error {
	stmt, err := db.db.Prepare("UPDATE goods SET location=?,price=?,restvolume=?,totalvolume=?,picture=?,uid=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gu.Location, gu.Price, gu.RestVolume, gu.TotalVolume, gu.Picture, gu.UID, gu.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Mysql) DeleteGoods(goodsID int64, ctx context.Context) error {
	stmt, err := db.db.Prepare("DELETE FROM goods WHERE id=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(goodsID)
	if err != nil {
		return err
	}

	affect, _ := res.RowsAffected()
	log.Infof("delete %v from db", affect)
	return nil
}

func (db *Mysql) Status(id int64) (int64, error) {
	stmt, err := db.db.Prepare("SELECT status FROM rule WHERE id=?")
	if err != nil {
		return 0, err
	}
	row := stmt.QueryRow(id)
	var status int64
	err = row.Scan(&status)
	if err != nil {
		return 0, err
	}
	return status, nil
}
