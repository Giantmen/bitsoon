package proto

type User struct {
	Id string
	Password string
	Phone string
	Email string
}

type UserInsert User
type UserInsertResult struct {
	//Id int64 //添加成功返回一个自增id作为自增id
}

type UserQueryAll struct {
}
type UserQueryAllResult []*User

type UserQueryOne struct {
	Id string //商品id
}
type UserQueryOneResult struct {
	Id    string
	Phone string
	Email string
}

type UserDelete struct {
	Id string
}

type UserUpdate struct {
	Id          string
	Password string
	Phone string
	Email string
}