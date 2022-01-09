package configs

import (
	"os"

	"rakshit.dev/m/src/utils"
)

var DGN string
var PORT string
var MongoURI string
var DBName string

func init() {
	DGN = "local"
	PORT = "80"
	MongoURI = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"
	DBName = "getir-case-study"
	env := os.Getenv("DEPLOYMENT_GROUP_NAME")
	if env != utils.EMPTY_STRING {
		DGN = env
	}
	env = os.Getenv("PORT")
	if env != utils.EMPTY_STRING {
		PORT = env
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
