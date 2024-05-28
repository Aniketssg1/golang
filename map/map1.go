package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	/* Using a for loop */
	// for key, value := range users {
	// 	if key == name && value.scheduledForDeleteion {
	// 		delete(users, key)
	// 		return true, nil
	// 	}
	// }
	// return false, errors.New("cannot be deleted")

	/* Check is name exists in map */
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if existingUser.scheduledForDeleteion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                  string
	number                int
	scheduledForDeleteion bool
}

func deleteUser() {
	users := map[string]user{
		"Aniket": {
			name:                  "Aniket",
			number:                123,
			scheduledForDeleteion: true,
		},
		"Abhinav": {
			name:                  "Abhinav",
			number:                324,
			scheduledForDeleteion: false,
		},
		"Sourabh": {
			name:                  "Sourabh",
			number:                657,
			scheduledForDeleteion: true,
		},
		"Ashwini": {
			name:                  "Ashwini",
			number:                76,
			scheduledForDeleteion: false,
		},
		"Sarthak": {
			name:                  "Sarthak",
			number:                324,
			scheduledForDeleteion: true,
		},
	}
	for key, _ := range users {
		deleted, err := deleteIfNecessary(users, key)
		fmt.Printf("name: %s deleted: %v error: %v\n", key, deleted, err)
	}
}
