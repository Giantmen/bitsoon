package store

import (
	"context"
	"fmt"

	"github.com/Giantmen/bitsoon/config"
	"github.com/Giantmen/bitsoon/proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"

)

type Mysql struct {
	dbObj  orm.Ormer
}

func init() {
	orm.RegisterModel(new(proto.User), new(proto.Goods))//goods 如果没有会自动创建
}

func NewMysql(cfg *config.Mysql) (*Mysql, error) {
	conn := cfg.ConnStr
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册mysql驱动
	//orm.RegisterModel(new(proto.Goods))
	//orm.RegisterModel(new(proto.User), new(proto.Goods))//goods 如果没有会自动创建
	//orm.RegisterDriver("mysql", orm.DRMySQL) //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", conn) //设置conn中的数据库为默认使用数据库
	orm.RunSyncdb("default", false, false)//后一个使用true会带上很多打印信息，数据库操作和建表操作的
	orm.Debug = true //true 打印数据库操作日志信息
	dbObj := orm.NewOrm()  //实例化数据库操作对象
	return &Mysql{
		dbObj:dbObj,
	},nil
}

func (orm *Mysql) Query(sql string,result interface{},ctx context.Context) (int64, error) {
	num, err := orm.dbObj.Raw(sql).QueryRows(result)
	if err != nil {
		fmt.Println("查询出错")
		return 0,err
	} else {
		fmt.Printf("共查询到记录:%d条\n", num)
		return num,nil
	}
}

func (orm *Mysql) Exec(sql string, ctx context.Context) (int64,error) {
	res, err := orm.dbObj.Raw(sql).Exec()
	if err != nil {
		fmt.Println("执行出错")
		return 0,err
	}
	id,_:= res.LastInsertId()
	return id,nil
}

