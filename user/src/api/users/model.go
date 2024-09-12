package users

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Email string             `json:"email,omitempty"`
	Age   int                `json:"age,omitempty"`
}

func (user User) ValidateEmail() (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(user.Email)
	return matchedString
}
func (u User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("firstname should not be empty")
	}
	if u.Email == "" || !u.ValidateEmail() {
		return fmt.Errorf("lastname should not be empty")
	}
	if u.Age <= 0 {
		return fmt.Errorf("address should not be empty")
	}
	return nil
}
