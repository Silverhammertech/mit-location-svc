package mongodb

import (
	"fmt"
	"github.com/Silverhammertech/mit-location-svc/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"strings"
)

func GetNear(resp *model.ResponseGeo) (err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %s", err.Error(), connectionString))
	}
	defer session.Close()

	// get sort order
	order := 1
	if strings.EqualFold(resp.Summary.Params.SortOrder, model.SORT_ORDER_DESC) {
		order = -1
	}

	// query the database
	c := session.DB(database).C(collection)
	err = c.Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{resp.Summary.Params.Longitude, resp.Summary.Params.Latitude}},
				"distanceField": "distance",
				"maxDistance":   resp.Summary.Params.Radius,
				"query":         "{}",
				"limit":         resp.Summary.Params.Limit,
				"spherical":     true,
			},
		},
		{
			"$sort": bson.M{resp.Summary.Params.SortParam: order},
		},
	},
	).All(&resp.Locations)
	return
}
