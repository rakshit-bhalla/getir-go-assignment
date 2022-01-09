package configs

import (
	"os"

	"rakshit.dev/m/src/utils"
)

var DGN string
var APIHost string
var MongoURI string
var DBName string

func init() {
	DGN = "local"
	APIHost = "localhost:8080"
	MongoURI = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"
	DBName = "getir-case-study"
	env := os.Getenv("DEPLOYMENT_GROUP_NAME")
	if env != utils.EMPTY_STRING {
		DGN = env
	}
	env = os.Getenv("API_HOST")
	if env != utils.EMPTY_STRING {
		APIHost = env
	}
	env = os.Getenv("MONGO_URI")
	if env != utils.EMPTY_STRING {
		MongoURI = env
	}
	env = os.Getenv("MONGO_DB_NAME")
	if env != utils.EMPTY_STRING {
		DBName = env
	}
}
