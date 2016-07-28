package util

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

func mainMongo() {
	//s := MongoInit()

	//query := bson.M{}
	//m, _ := Find(s, "weixin", "user", query)
}

//================================================================================
func MongoInit(host string) *mgo.Session {
	session, err := mgo.Dial(host)
	Check(err)
	return session
}

func execute(s *mgo.Session, db, tb string, do func(c *mgo.Collection) error) error {
	coll := s.DB(db).C(tb)
	err := do(coll)
	return err
}

func Insert(s *mgo.Session, db, col string, o interface{}) error {
	queryFunc := func(c *mgo.Collection) error {
		return c.Insert(o)
	}
	err := execute(s, db, col, queryFunc)
	return err
}

func Find(s *mgo.Session, db, col string, query bson.M) ([]map[string]interface{}, error) {
	m := []map[string]interface{}{}
	queryFunc := func(c *mgo.Collection) error {
		return c.Find(query).All(&m)
	}
	err := execute(s, db, col, queryFunc)
	return m, err
}

func Update(s *mgo.Session, db, col string, query, update bson.M) error {
	queryFunc := func(c *mgo.Collection) error {
		_, err := c.UpdateAll(query, update)
		return err
	}
	err := execute(s, db, col, queryFunc)
	return err
}

func DeleteSoft(s *mgo.Session, db, col string, query bson.M) error {
	queryFunc := func(c *mgo.Collection) error {
		_, err := c.UpdateAll(query, bson.M{"$set": bson.M{"delete": true, "uptime": time.Now().Unix()}})
		return err
	}
	err := execute(s, db, col, queryFunc)
	return err
}

func DeleteHard(s *mgo.Session, db, col string, query bson.M) error {
	queryFunc := func(c *mgo.Collection) error {
		return c.Remove(query)
	}
	err := execute(s, db, col, queryFunc)
	return err
}

func Count(s *mgo.Session, db, col string, query bson.M) (int, error) {
	var n int
	queryFunc := func(c *mgo.Collection) error {
		var err error
		n, err = c.Find(query).Count()
		return err
	}
	err := execute(s, db, col, queryFunc)
	return n, err
}
