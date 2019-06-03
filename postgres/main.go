package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB

func init() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5433",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		PoolSize: 20,
	})
	//defer db.Close()
}

func main() {
	ExampleDB_Model()
	//ExampleDB_Insert_dynamicTableName()
	//time.Sleep(15 * time.Second)
}

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func ExampleDB_Model() {

	//err := createSchema(db)
	//if err != nil {
	//	panic(err)
	//}

	var err error

	//user1 := &User{
	//	Name:   "admin",
	//	Emails: []string{"admin1@admin", "admin2@admin"},
	//}
	//err = db.Insert(user1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.Insert(&User{
	//	Name:   "root",
	//	Emails: []string{"root1@root", "root2@root"},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//story1 := &Story{
	//	Title:    "Cool story",
	//	AuthorId: user1.Id,
	//}
	//err = db.Insert(story1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Select user by primary key.
	//user := &User{Id: user1.Id}

	//user := &User{Id: 5}
	//err = db.Select(user)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(user)

	t1 := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			user := &User{Id: 1}
			err = db.Select(user)
			if err != nil {
				panic(err)
			}
			if i%1000 == 0 {
				fmt.Println(i, user)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	//
	// Select all users.
	//var users []User
	//err = db.Model(&users).Where("name=?", "root").Select()
	//if err != nil {
	//	panic(err)
	//}

	// Select story and associated author in one query.
	//story := new(Story)
	//err = db.Model(story).
	//	Relation("Author").
	//	Where("story.id = ?", 1).
	//	Select()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(user)
	//fmt.Println(users)
	//fmt.Println(story)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
		fmt.Println(model)
		err := db.CreateTable(model, &orm.CreateTableOptions{
			//Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ExampleDB_Insert_dynamicTableName() {
	type NamelessModel struct {
		tableName struct{} `sql:"_"` // "_" means no name
		Id        int
	}

	err := db.Model((*NamelessModel)(nil)).Table("dynamic_name").CreateTable(nil)
	panicIf(err)

	row123 := &NamelessModel{
		Id: 123,
	}
	_, err = db.Model(row123).Table("dynamic_name").Insert()
	panicIf(err)

	row := new(NamelessModel)
	err = db.Model(row).Table("dynamic_name").First()
	panicIf(err)
	fmt.Println("id is", row.Id)

	err = db.Model((*NamelessModel)(nil)).Table("dynamic_name").DropTable(nil)
	panicIf(err)

	// Output: id is 123
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
