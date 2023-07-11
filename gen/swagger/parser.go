package swagger

import (
	"github.com/iancoleman/strcase"
	"io"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"strings"
)

type Parameter struct {
	Name        string
	Description string
	Required    bool
}

type Model struct {
	Name       string
	Parameters []Parameter
}

type API struct {
	Name        string
	Description string
	Parameters  []Parameter
}

type Swagger struct {
	obj interface{}
}

func New(fileUrl string) (*Swagger, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}

	swaggerYaml, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = yaml.Unmarshal(swaggerYaml, &v)
	if err != nil {
		panic(err)
	}

	return &Swagger{obj: v}, nil
}

func (s *Swagger) GetModel(modelName string) *Model {
	if root, ok := s.obj.(map[string]interface{}); ok {
		if models, ok := root["definitions"].(map[string]interface{}); ok {
			for candidate, fields := range models {
				snakeModelName := strcase.ToSnake(modelName)
				snakeModelName = strings.ReplaceAll(snakeModelName, "apiaws", "api_aws")
				snakeModel := strcase.ToSnake(strings.ReplaceAll(candidate, ".", "_"))
				idx := 45
				if len(snakeModel) < idx {
					idx = len(snakeModel)
				}
				if len(snakeModelName) < idx {
					idx = len(snakeModelName)
				}

				if snakeModel == snakeModelName {
					modelFields := fields.(map[string]interface{})
					requireds := modelFields["required"]
					model := Model{Name: candidate}

					descriptionMap := map[string]string{}
					properties := modelFields["properties"]
					if properties != nil {
						propertiesMap := properties.(map[string]interface{})
						for key, prop := range propertiesMap {
							propMap := prop.(map[string]interface{})
							if propMap["description"] != nil {
								descriptionMap[key] = propMap["description"].(string)
							}
						}
					}

					if requireds != nil {
						requiredArray := requireds.([]interface{})
						for _, param := range requiredArray {
							paramName := param.(string)
							model.Parameters = append(model.Parameters, Parameter{
								Name:        paramName,
								Description: descriptionMap[paramName],
								Required:    true,
							})
						}
					}

					return &model
				}
			}
		}
	}
	return nil
}

func (s *Swagger) GetAPI(modelName string) *API {
	root := s.obj.(map[string]interface{})
	models := root["paths"].(map[string]interface{})
	for candidate, children := range models {
		methodMap := children.(map[string]interface{})
		for method, reqs := range methodMap {
			apiName := method + "_" + strings.ReplaceAll(candidate, "/", "_")
			apiName = strings.ReplaceAll(apiName, "{", "_")
			apiName = strings.ReplaceAll(apiName, "}", "_")
			apiName = strcase.ToCamel(apiName) + "Params"
			apiName = strings.ReplaceAll(apiName, "ApiV", "APIV")
			apiName = strings.ReplaceAll(apiName, "Id", "ID")

			if apiName == modelName {
				reqsMap := reqs.(map[string]interface{})
				var params []interface{}
				if reqsMap["parameters"] != nil {
					params = reqsMap["parameters"].([]interface{})
				}

				apiDescription, apiDescriptionExists := reqsMap["description"]
				var apiDesc string
				if str, ok := apiDescription.(string); ok && apiDescriptionExists {
					apiDesc = str
				}

				api := &API{Name: candidate, Description: apiDesc}
				for _, param := range params {
					paramMap := param.(map[string]interface{})
					required, requiredExists := paramMap["required"]

					description, descriptionExists := paramMap["description"]
					var desc string
					if str, ok := description.(string); ok && descriptionExists {
						desc = str
					}

					api.Parameters = append(api.Parameters, Parameter{
						Name:        paramMap["name"].(string),
						Required:    requiredExists && required.(bool),
						Description: desc,
					})
				}
				return api
			}
		}
	}
	return nil
}
