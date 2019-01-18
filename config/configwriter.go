package config

import (
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

// Client clients to show the results
type Client struct {
	ID    int    `json:"id"`
	Place string `json:"place"`
	URL   string `json:"url"`
}

// ChangeOrAddClient if the client by ID already exists a new client will be added
func (r WriterSettings) ChangeOrAddClient(client *Client) {

	type Clients struct {
		Key []Client `json:"clients"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open(r.FilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonParsed, err := gabs.ParseJSON([]byte(byteValue))
	children, _ := jsonParsed.S("clients").Children()
	fmt.Println(jsonParsed.Path("clients.id").String())
	idExists := false
	index := 0
	for _, child := range children {
		yes := child.Exists("id")
		if yes == true {
			fmt.Println("EXIST")
		}
		//fmt.Println(child.Path("id").String())
		value, error := strconv.Atoi(child.Search("id").String())
		if error == nil {

			if value == client.ID {
				fmt.Println("HI")
				idExists = true
				break
			}
		}
		index = index + 1
	}

	if idExists {
		fmt.Println("REMOVE " + strconv.Itoa(index))
		jsonParsed.ArrayRemove(index, "clients")
	}

	jsonParsed.ArrayAppend(client, "clients")

	jsonParsed.StringIndent("", "  ")

	prettyJSON := jsonParsed.StringIndent("", "  ")
	fmt.Println(prettyJSON)
	//jsonByte, _ := json.Marshal(prettyJSON)
	//err = ioutil.WriteFile(r.FilePath, jsonByte, 0644)
}
