package repository

import (
	"github.com/juju/mgosession"
	"github.com/rsegura/cocinario-api.git/pkg/foods"
	bson "gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "foods"
)

type mongoFoodsRepository struct {
	Conn   *mgosession.Pool
	dbName string
}

func NewMongoFoodsRepository(p *mgosession.Pool, dbName string) foods.Repository {
	return &mongoFoodsRepository{
		Conn:   p,
		dbName: dbName,
	}
}

func (m *mongoFoodsRepository) GetById(id string) (interface{}, error) {
	bsonId := bson.ObjectIdHex(id)
	session := m.Conn.Session(nil)
	coll := session.DB(m.dbName).C(COLLECTION)
	result := bson.M{}
	err := coll.Find(bson.M{"_id": bsonId}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *mongoFoodsRepository) Fetch() (interface{}, error) {
	session := m.Conn.Session(nil)
	coll := session.DB(m.dbName).C(COLLECTION)
	var result = make([]interface{}, 0, 10)
	err := coll.Find(nil).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
