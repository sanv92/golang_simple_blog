package json

import (
	"encoding/json"
	"os"
	//"fmt"
)

/*
func getData() ([]*News, error) {
	file, err := os.Open(menuFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*News
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
*/

/*
func ReadJSON() ([]*Menu, error) {
	// Open the file.
	file, err := os.Open(MenuFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var menu []*Menu
	fmt.Println("menu: ", menu)

	err = json.NewDecoder(file).Decode(&menu)

	// We don't need to check for errors, the caller can do this.
	return menu, err
}
*/
/*
func ReadJSON(result interface{}) error {
	fmt.Println("result: ", result)
	// Open the file.
	file, err := os.Open(MenuFile)
	if err != nil {
		return err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	err = json.NewDecoder(file).Decode(&result)

	// We don't need to check for errors, the caller can do this.
	return err
}*/

func ReadJSON(result interface{}) error {
	file, err := os.Open(MenuFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(result)
}
