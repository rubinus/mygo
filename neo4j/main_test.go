package main

import (
	"fmt"
	"testing"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

/*
@Time : 2019-06-03 14:10
@Author : rubinus.chu
@File : main_test
@project: mygo
*/

func getPeople(driver neo4j.Driver) ([]string, error) {
	var people interface{}
	var err error
	var session neo4j.Session

	// 获取neo4j session, 切记关闭session
	if session, err = driver.Session(neo4j.AccessModeRead); err != nil {
		return nil, err
	}
	defer session.Close()

	people, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string
		var result neo4j.Result

		if result, err = tx.Run("MATCH (a:Person) RETURN a.name ORDER BY a.name", nil); err != nil {
			return nil, err
		}

		for result.Next() {
			// 输出结果集中的记录
			fmt.Println(result.Record())
			list = append(list, result.Record().GetByIndex(0).(string))
		}

		if err = result.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})
	if err != nil {
		return nil, err
	}

	return people.([]string), nil
}

// Neo4j-测试获取结果集
func TestGetPeople(t *testing.T) {
	var driver neo4j.Driver
	var err error
	var result []string

	// 创建neo4j驱动
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	if err != nil {
		panic("Get Driver Failed: " + err.Error())
	}
	defer driver.Close()

	// 获取结果
	result, err = getPeople(driver)
	if err != nil {
		panic("Get People Failed: " + err.Error())
	}

	fmt.Println(result)
}

func TestTestNeo4j(t *testing.T) {
	TestNeo4j()
}
