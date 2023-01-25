package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	State   string `json:"state" bson:"user_state"`
	City    string `json:"city" bson:"user_city"`
	Pincode int    `json:"pincode" bson:"user_pincode"`
}

type User struct {
	// ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"user_name"`
	Age         int                `json:"age" bson:"user_age"`
	UserAddress Address            `json:"address" bson:"user_address"`
}
