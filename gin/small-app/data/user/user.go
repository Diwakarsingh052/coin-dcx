package user

import "errors"

var users = map[uint64]User{
	123: {
		FName: "Bob",
		LName: "abc",
		Email: "bob@email.com",
	},
}
var ErrUserNotFound = errors.New("user id not found")

// FetchUser fetches the data from the map
func FetchUser(id uint64) (User, error) {
	u, ok := users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}
