package connector

import backend "ashish.com/backend/database"

func Connector() string {

	backend.Database()
	return "CONNECTOR\n"
}
