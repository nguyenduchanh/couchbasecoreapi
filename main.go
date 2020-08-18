package main

import (
	"couchbasecoreapi/router"
	_ "io/ioutil"
)

//api for couchbase management
//tam thoi load mac dinh server cuar minh de thanh hinh
//sau do se load tham so tu file json
func main() {
	router.Router()
}
