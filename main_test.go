package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/garyburd/redigo/redis"
	mgo "gopkg.in/mgo.v2"
)

func TestMain(t *testing.T) {
	main()
}

func TestHello(t *testing.T) {
	result := hello()
	if "hello" != result {
		t.Error("it's not hello")
	}
}

func TestRedis(t *testing.T) {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	t.Log(c.Do("PING"))
}

func TestMongo(t *testing.T) {
	url := "mongo"
	session, err := mgo.Dial(url)
	if err != nil {
		t.Error(err)
	}
	t.Log(session.Ping())
}

func TestMaria(t *testing.T) {
	db, err := sql.Open("mysql", "root:test@(mariadb:3306)/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	t.Log(db.Ping())
}
