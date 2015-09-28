package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

// Model Struct
// type Article struct {
//   Article_ID int
//   Title string `orm:"type(text)"`
//   Content string `orm:"type(text)"`
//   Author string `orm:"type(text)"`
//   Created_at string `orm:"size(10)"`
//   Updated_at string `orm:"size(10)"`
// }

func init() {
    // set default database
    //orm.RegisterDataBase("default", "mysql", "root:root@/br_rigo?charset=utf8", 30)
    orm.RegisterDataBase("default", "mysql", "root:root@/br_rigo?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()

    user := User{Name: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("ERR: %v\n", err)

    // delete
    // num, err = o.Delete(&u)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
