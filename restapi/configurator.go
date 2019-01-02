package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/middleware/untyped"
)

// NewConfigurator creates a new configuration api handler
func NewConfigurator() (http.Handler, error) {
	yamlFile, err := os.Open("./restapi/swaggerui/configurator.yml")
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()
	yamlFileData, err := ioutil.ReadAll(yamlFile)

	json, _ := yaml.YAMLToJSON(yamlFileData)

	spec, err := loads.Analyzed(json, "")
	if err != nil {
		return nil, err
	}
	api := untyped.NewAPI(spec)

	api.RegisterOperation("get", "/hello", getGreeting)

	return middleware.Serve(spec, api), nil
}

// AddSwaggerUIToHandler creates a swagger ui api handler
func AddSwaggerUIToHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
		} else {
			handler := http.FileServer(http.Dir("./restapi/swaggerui"))
			handler = http.StripPrefix("/swaggerui/", handler)
			handler.ServeHTTP(w, r)
		}
	})
}

var getGreeting = runtime.OperationHandlerFunc(func(data interface{}) (interface{}, error) {
	name := data.(map[string]interface{})["name"].(string)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, nil
})
