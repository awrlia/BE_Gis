package peda

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunanLineString(mconn, collectionname)
	return GCFReturnStruct(datagedung)
}

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func CreateUser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		hash, hashErr := HashPassword(datauser.Password)
		if hashErr != nil {
			response.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(mconn, collname, datauser.Username, datauser.Role, hash)
		response.Message = "Berhasil Input data"
	}
	return GCFReturnStruct(response)
}
func MembuatGeojsonPointToken(mongoenv, dbname, collname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(mongoenv, dbname)

	// Parsing Request Body
	var datapoint GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&datapoint)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		// Handling Authorization
		err := PostLinestring(mconn, collname, datapoint)
		if err != nil {
			// Success
			return GCFReturnStruct(CreateResponse(true, "Success: LineString created", datapoint))
		} else {
			return GCFReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// This part is unreachable, so you might want to remove it
	// return GCFReturnStruct(CreateResponse(false, "Success to create LineString", nil))
}

func MembuatGeojsonPolylineToken(mongoenv, dbname, collname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(mongoenv, dbname)

	// Parsing Request Body
	var datapoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&datapoint)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		// Handling Authorization
		err := PostPoint(mconn, collname, datapoint)
		if err != nil {
			// Success
			return GCFReturnStruct(CreateResponse(true, "Success: LineString created", datapoint))
		} else {
			return GCFReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// This part is unreachable, so you might want to remove it
	// return GCFReturnStruct(CreateResponse(false, "Success to create LineString", nil))
}

func MembuatGeojsonPoligonToken(mongoenv, dbname, collname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(mongoenv, dbname)

	// Parsing Request Body
	var datapoint GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&datapoint)
	if err != nil {
		return err.Error()
	}

	if r.Header.Get("token") == os.Getenv("token") {
		// Handling Authorization
		err := PostPolygon(mconn, collname, datapoint)
		if err != nil {
			// Success
			return GCFReturnStruct(CreateResponse(true, "Success: LineString created", datapoint))
		} else {
			return GCFReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// This part is unreachable, so you might want to remove it
	// return GCFReturnStruct(CreateResponse(false, "Success to create LineString", nil))
}

func MengambilGeojsonToken(mongoenv, dbname, collname string, r *http.Request) string {
	// MongoDB Connection Setup
	mconn := SetConnection(mongoenv, dbname)

	if r.Header.Get("token") == os.Getenv("token") {
		// Handling Authorization
		datagedung := GetAllBangunanLineString(mconn, collname)
		err := json.NewDecoder(r.Body).Decode(&datagedung)
		if err != nil {
			// Success
			return GCFReturnStruct(CreateResponse(true, "Success: LineString created", datagedung))
		} else {
			return GCFReturnStruct(CreateResponse(false, "Error", nil))
		}
	} else {
		return GCFReturnStruct(CreateResponse(false, "Unauthorized: Secret header does not match", nil))
	}

	// This part is unreachable, so you might want to remove it
	// return GCFReturnStruct(CreateResponse(false, "Success to create LineString", nil))
}