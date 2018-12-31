package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ghodss/yaml"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/middleware/untyped"
)

// NewPetstore creates a new petstore api handler
func NewConfigurator() (http.Handler, error) {
	yamlFile, err := os.Open("./ui/configurator.yml")
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()
	yamlFileData, err := ioutil.ReadAll(yamlFile)

	//fmt.Println(yamlFileData)

	json, _ := yaml.YAMLToJSON(yamlFileData)

	spec, err := loads.Analyzed(json, "")
	if err != nil {
		return nil, err
	}
	api := untyped.NewAPI(spec)

	api.RegisterOperation("get", "/hello", getGreeting)

	return middleware.Serve(spec, api), nil
}

var getGreeting = runtime.OperationHandlerFunc(func(data interface{}) (interface{}, error) {
	fmt.Println("getPetByID")
	name := data.(map[string]interface{})["name"].(string)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, nil
})
