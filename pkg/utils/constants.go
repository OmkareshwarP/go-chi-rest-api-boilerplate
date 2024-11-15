package utils

import (
	"go-chi-rest-api-boilerplate/pkg/types"
	"regexp"
	"sync"
)

var (
	UsersData = []*types.User{} // in-memory store
	Mu        sync.Mutex        // protects usersData
)

// Helper function to check if a username is unique
func IsUsernameUnique(userId string, username string) bool {
	for _, user := range UsersData {
		if user.UserID != userId && user.Username == username {
			return false
		}
	}
	return true
}

func RemoveSpacesAndSpecialChars(input string) string {
	// Create a regular expression to match non-alphanumeric characters
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	// Replace matched characters with an empty string
	return re.ReplaceAllString(input, "")
}
