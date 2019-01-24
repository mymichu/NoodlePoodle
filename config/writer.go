package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Jeffail/gabs"
)

// WriterSettings bla
type WriterSettings struct {
	FilePath string
}

func (r WriterSettings) AddClient(client *Client) error {
	// Open our jsonFile
	//jsonFile, err := os.Open(r.FilePath)
	// if we os.Open returns an error then handle it
	//if err != nil {
	//	fmt.Println(err)
	//}
	return nil
}

// ChangeClient if the client by ID already exists the client will be modified
func (r WriterSettings) ChangeClient(client *Client) error {

	// Open our jsonFile
	jsonFile, err := os.Open(r.FilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return errors.New("Can't open config-file: " + err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonParsed, err := gabs.ParseJSON([]byte(byteValue))
	children, _ := jsonParsed.S("clients").Children()
	fmt.Println(jsonParsed.Path("clients.id").String())

	index, idExist := checkIfIDExist(children, uint16(client.ID))

	if !idExist {
		return errors.New("ID " + string(client.ID) + " does not exist")
	}
	//Remove current id and Add
	jsonParsed.ArrayRemove(int(index), "clients")
	jsonParsed.ArrayAppend(client, "clients")
	r.writeJSONToFile(jsonParsed)

	return nil
}

func checkIfIDExist(children []*gabs.Container, id uint16) (uint16, bool) {
	var index uint16
	var exist bool
	for _, child := range children {
		idExist := child.Exists("id")
		if idExist == true {
			value, error := strconv.ParseInt(child.Search("id").String(), 10, 16)
			if error == nil {
				if uint16(value) == id {
					exist = true
					break
				}
			}
		}
		index = index + 1
	}
	return index, exist
}

func (r WriterSettings) writeJSONToFile(jsonContainer *gabs.Container) {
	jsonContainer.StringIndent("", "  ")

	prettyJSON := jsonContainer.StringIndent("", "  ")
	fmt.Println(prettyJSON)
	//jsonByte, _ := json.Marshal(prettyJSON)
	//err = ioutil.WriteFile(r.FilePath, jsonByte, 0644)
}
