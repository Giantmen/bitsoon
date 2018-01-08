# bitsoon
1.	用户操作
λ	添加用户
path: /user/insert
method:post
参数：
{
	"name":"linc",
	"password":"lc",
	"phone":"183",
	"email":"sdfa@163.com"
}
返回值：
   {
    "Code": 200,
    "Data": 2,
    "Msg": "success"
}
data域为新增用户的自增id

λ	删除用户
path:/user/delete
method:post
参数：
{
	"id":1
}
返回值：
{
    "Code": 200,
    "Data": null,
    "Msg": "success"
}
λ	更新用户
path:/user/update
method:post
参数：
{
	"id":2,
	"name":"linc",
	"password":"lc",
	"phone":"184",
	"email":"sdfa@163.com"
}
返回值：
{
    "Code": 200,
    "Data": null,
    "Msg": "success"
}
查询用户信息：
path:/user/query/{userID}
method:get
返回值：
{
    "Code": 200,
    "Data": {
        "Id": 2,
        "Name": "linc",
        "Password": "lc",
        "Phone": "183",
        "Email": "sdfa@163.com"
    },
    "Msg": "success"
}

2．	商品操作
λ	添加商品
path: /user/insert
method:post
参数：
{
	"name":"linc",
	"password":"lc",
	"phone":"183",
	"email":"sdfa@163.com"
}
返回值：
   {
    "Code": 200,
    "Data": 2,
    "Msg": "success"
}
data域为新增商品的自增id

λ	删除用户
path:/user/delete
method:post
参数：
{
	"id":1
}
返回值：
{
    "Code": 200,
    "Data": null,
    "Msg": "success"
}
λ	更新用户
path:/user/update
method:post
参数：
{
	"id":2,
	"name":"linc",
	"password":"lc",
	"phone":"184",
	"email":"sdfa@163.com"
}
返回值：
{
    "Code": 200,
    "Data": null,
    "Msg": "success"
}
查询用户信息：
path:/user/query/{userID}
method:get
返回值：
{
    "Code": 200,
    "Data": {
        "Id": 2,
        "Name": "linc",
        "Password": "lc",
        "Phone": "183",
        "Email": "sdfa@163.com"
    },
    "Msg": "success"
}

