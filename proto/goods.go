package proto

type Goods struct {
	ID          int64
	Location    string
	Price       float64
	RestVolume  int
	TotalVolume int
	Picture     string
	UID         string
}

type GoodsInsert struct {
	//ID          int64
	Location    string
	Price       float64
	RestVolume  int
	TotalVolume int
	Picture     string
	UID         string
}
type GoodsInsertResult struct {
	ID int64 //添加成功返回一个自增id作为自增id
}

type GoodsQueryAll struct {
}
type GoodsQueryAllResult struct {
	Gs []*Goods
}

type GoodsQueryOne struct {
	ID int64 //商品id
}
type GoodsQueryOneResult struct {
	ID          int64
	Location    string
	Price       float64
	RestVolume  int
	TotalVolume int
	Picture     string
	UID         string
}

type GoodsDelete struct {
	ID int64
}

type GoodsUpdate struct {
	ID          int64
	Location    string
	Price       float64
	RestVolume  int
	TotalVolume int
	Picture     string
	UID         string
}

type Response struct {
	Code int
	Data interface{}
	Msg  string
}
