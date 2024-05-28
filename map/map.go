package main

import "fmt"

func main() {
	person := map[string]string{
		"Name":  "Aniket",
		"Place": "Pune",
	}

	place := make(map[string]string)
	place["Name"] = "Pune"
	place["Country"] = "India"

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println(place)

	websites := map[string]map[string]string{
		"Google": map[string]string{
			"name":    "Google",
			"website": "https://www.google.com",
		},
		"MKCL": map[string]string{
			"name":    "MKCL",
			"website": "https://www.mkcl.org/",
		},
	}

	for key, value := range websites {
		fmt.Print(key)
		for name, website := range value {
			fmt.Printf("%s : %s\n", name, website)
		}
	}

	deleteUser()
}
