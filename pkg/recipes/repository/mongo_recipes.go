package repository

import (
	"github.com/juju/mgosession"
	"github.com/rsegura/cocinario-api.git/pkg/recipes"
)

const (
	COLLECTION = "recipes"
)

type mongoRecipesRepository struct {
	Conn   *mgosession.Pool
	dbName string
}

func NewMongoRecipesRepository(p *mgosession.Pool, dbName string) recipes.Repository {
	return &mongoRecipesRepository{
		Conn:   p,
		dbName: dbName,
	}
}

func (m *mongoRecipesRepository) Fetch() (interface{}, error) {
	session := m.Conn.Session(nil)
	coll := session.DB(m.dbName).C(COLLECTION)
	var result = make([]interface{}, 0, 10)
	err := coll.Find(nil).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
