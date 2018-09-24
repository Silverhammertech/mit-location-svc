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

func GetOtherProducts() (products []model.OtherProduct, err error) {

	prod, err := GetEdibleProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetConcentrateProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	return
}

func GetEdibleProducts() (products []model.OtherProduct, err error) {

	prod, err := GetEdibleBeverageProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleBreakFastProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleBrownieProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleCandyProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleCapsuleProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleChocolateProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleCondimentProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleCookieProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleCookingProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleFrozenProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleSnackProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetEdibleTinctureProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	return
}

func GetEdibleBeverageProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_beverage")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleBreakFastProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_breakfast")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleBrownieProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_brownie")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleCandyProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_candy")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleCapsuleProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_capsule")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleChocolateProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_chocolate")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleCondimentProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_condiment")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleCookieProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_cookie")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleCookingProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_cooking")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleFrozenProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_frozen")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleSnackProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_snack")
	err = c.Find(nil).All(&products)

	return
}

func GetEdibleTinctureProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("edible_tincture")
	err = c.Find(nil).All(&products)

	return
}

func GetConcentrateProducts() (products []model.OtherProduct, err error) {

	prod, err := GetConcentrateCartridgeProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetConcentrateIngestibleProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetConcentrateSolventProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetConcentrateSolventlessProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	prod, err = GetConcentrateTerpeneProducts()
	if err != nil {
		return []model.OtherProduct{}, err
	}
	products = append(products, prod...)

	return
}

func GetConcentrateCartridgeProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("concentrate_cartridge")
	err = c.Find(nil).All(&products)

	return
}

func GetConcentrateIngestibleProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("concentrate_ingestible")
	err = c.Find(nil).All(&products)

	return
}

func GetConcentrateSolventProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("concentrate_solvent")
	err = c.Find(nil).All(&products)

	return
}

func GetConcentrateSolventlessProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("concentrate_solventless")
	err = c.Find(nil).All(&products)

	return
}

func GetConcentrateTerpeneProducts() (products []model.OtherProduct, err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return []model.OtherProduct{}, errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C("concentrate_terpene")
	err = c.Find(nil).All(&products)

	return
}
