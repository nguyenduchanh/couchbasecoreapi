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

//func Example() {
//	cluster, _ := gocb.Connect("http://localhost")
//	cluster.Authenticate(gocb.PasswordAuthenticator{
//		Username: "Administrator",
//		Password: "abc123",
//	})
//
//	bucket, _ := cluster.OpenBucket("classroom", "")
//
//	myQuery := gocb.NewN1qlQuery("SELECT name, age FROM `classroom` ")
//	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
//	if err != nil {
//		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
//	}
//
//	// Iterate through rows and print output
//	var row model.Query
//	var retValues []model.Query
//
//	// Stream the values returned from the query into a typed array of structs
//	for rows.Next(&row) {
//
//		// Check if the current row has a value for FAA, if it does it's an airport
//		//  and should be added to the return values
//		retValues = append(retValues, row)
//
//		// Set the row to an empty struct, to prevent current values being added
//		//  to the next row in the results collection returned by the query
//		row = model.Query{}
//	}
//
//	// Marshal array of structs to JSON
//	bytes, err := json.Marshal(retValues)
//	if err != nil {
//		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
//	}
//	//var t []model.Query
//	json.Unmarshal(bytes, retValues)
//	fmt.Print(retValues)
//	// Exiting
//}
