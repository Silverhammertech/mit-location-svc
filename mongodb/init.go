package mongodb

import "os"

var connectionString string
var database string
var collection string

func init() {
	connectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		// for local
		connectionString = "mongodb://admin:Passw0rd@ds159631.mlab.com:59631/silverhammer"
	}

	database = os.Getenv("MONGODB_DB")
	if len(database) == 0 {
		// for local
		database = "silverhammer"
	}

	collection = os.Getenv("MONGODB_COLLECTION")
	if len(collection) == 0 {
		// for local
		collection = "mitlocation"
	}
}
