package proto

type Goods struct {
	Id          int64
	Location    string
	Price       float64
	Restvolume  int
	Totalvolume int
	Picture     string
	Uid         string
}

type GoodsInsert struct {
	//ID          int64
	Location    string
	Price       float64
	Restvolume  int
	Totalvolume int
	Picture     string
	Uid         string
}
type GoodsInsertResult struct {
	Id int64 //添加成功返回一个自增id作为自增id
}

type GoodsQueryAll struct {
}
type GoodsQueryAllResult []*Goods

type GoodsQueryOne struct {
	Id int64 //商品id
}
type GoodsQueryOneResult struct {
	Id          int64
	Location    string
	Price       float64
	RestVolume  int
	TotalVolume int
	Picture     string
	UID         string
}

type GoodsDelete struct {
	Id int64
}

type GoodsUpdate struct {
	Id          int64
	Location    string
	Price       float64
	Restvolume  int
	Totalvolume int
	Picture     string
	Uid         string
}

type Response struct {
	Code int
	Data interface{}
	Msg  string
}
