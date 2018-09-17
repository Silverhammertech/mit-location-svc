package mongodb

import (
	"fmt"
	"github.com/Silverhammertech/mit-location-svc/model"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

func GetAllProducts() (products []model.Product, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.Product{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()
	var tempProducts []model.Product

	// query the database
	c := session.DB(database).C("strain_hybrid")
	err = c.Find(nil).All(&tempProducts)
	products = append(products, tempProducts...)

	c = session.DB(database).C("strain_sativa")
	err = c.Find(nil).All(&tempProducts)
	products = append(products, tempProducts...)

	c = session.DB(database).C("strain_indica")
	err = c.Find(nil).All(&tempProducts)
	products = append(products, tempProducts...)

	return
}

func GetHybridProducts() (products []model.Product, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.Product{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("strain_hybrid")
	err = c.Find(nil).All(&products)

	return
}

func GetIndicaProducts() (products []model.Product, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.Product{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("strain_indica")
	err = c.Find(nil).All(&products)

	return
}

func GetSativaProducts() (products []model.Product, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.Product{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("strain_sativa")
	err = c.Find(nil).All(&products)

	return
}
