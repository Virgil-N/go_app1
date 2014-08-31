//注释的代码也可以正常运行的
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

type Test struct {
	//字段名字必须首字母大写
	Title string
	By    string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//选择数据库和集合
	c := session.DB("db1").C("test1")
	/*
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			panic(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			panic(err)
		}
	*/

	result := []Test{}
	err = c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	for i, length := 0, len(result); i < length; i++ {
		//字段名字必须首字母大写
		fmt.Println("title: ", result[i].Title, " by: ", result[i].By)
	}
	//fmt.Println("name: ", result[0].Title)
}
