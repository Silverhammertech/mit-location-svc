package mongodb

import (
	"fmt"
	"github.com/Silverhammertech/mit-location-svc/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

func GetState(resp *model.ResponseGeo) (err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// query the database
	c := session.DB(database).C(collection)
	err = c.Find(bson.M{"state": resp.Summary.Params.State}).Sort(resp.Summary.Params.SortParam).All(&resp.Locations)

	return
}
