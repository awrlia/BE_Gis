package peda

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []GeoJson {
	lokasi := atdb.GetAllDoc[[]GeoJson](mongoconn, collection)
	return lokasi
}

func PostPoint(mongoconn *mongo.Database, collection string, pointdata GeoJsonPoint) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, pointdata)
}

func PostLinestring(mongoconn *mongo.Database, collection string, linestringdata GeoJsonLineString) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, linestringdata)
}

func PostPolygon(mongoconn *mongo.Database, collection string, polygondata GeoJsonPolygon) interface{} {
	return atdb.InsertOneDoc(mongoconn, collection, polygondata)
}


func IsPasswordValid(mongoconn *mongo.Database, collection string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mongoconn, collection, filter)
	return CheckPasswordHash(userdata.Password, res.Password)
}

func InsertUserdata(mongoenv *mongo.Database, collname, username, role, password string) (InsertedID interface{}) {
	req := new(User)
	req.Username = username
	req.Password = password
	req.Role = role
	return atdb.InsertOneDoc(mongoenv, collname, req)
}

func CreateResponse(status bool, message string, data interface{}) Jaja {
	response := Jaja{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}