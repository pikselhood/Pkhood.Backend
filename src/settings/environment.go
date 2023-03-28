package settings

import "os"

var MongoConnection string
var DatabaseName = "pikseltb"
var CollectionName = "accounts"

func InitEnvironment() {
	var env = os.Getenv("env")

	switch env {
	case "dev":
		MongoConnection = dev_MongoConnection
	}
}
