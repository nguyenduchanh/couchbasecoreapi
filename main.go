package main

import (
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

//api for couchbase management
//tam thoi load mac dinh server cuar minh de thanh hinh
//sau do se load tham so tu file json
func main() {

}
func GetAllBucket() {

}
func Example() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "abc123",
	})

	bucket, _ := cluster.OpenBucket("classroom", "")

	bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	bucket.Upsert("u:kingarthur",
		User{
			Name: "kingarthur",
			Age:  "33",
		}, 0)

	// Get the value back
	var inUser User
	bucket.Get("u:kingarthur", &inUser)
	fmt.Printf("User: %v\n", inUser)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM classroom ")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"African Swallows"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}
}
