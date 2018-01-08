package proto

type User struct {
	Id int64    
	Name string 
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
	Id int64 //用户id
}
type UserQueryOneResult struct {
	Id    int64
	Name  string
	Phone string
	Email string
}

type UserDelete struct {
	Id int64
}

type UserUpdate struct {
	Id          int64
	Name     string
	Password string
	Phone string
	Email string
}
