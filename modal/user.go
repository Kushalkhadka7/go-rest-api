package modal

// User models the data type to be stored in db.
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	PhoneNo int    `json:"phoneNo"`
}
