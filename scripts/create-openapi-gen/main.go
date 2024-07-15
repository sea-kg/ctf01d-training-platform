package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

func dereferenceSchema(schema *spec.Schema) (*spec.Schema, error) {
	opts := &spec.ExpandOptions{
		RelativeBase: "openapi_dereferenced.yaml",
	}
	err := spec.ExpandSchemaWithBasePath(schema, nil, opts)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func main() {
	doc, err := loads.Spec("openapi.yaml")
	if err != nil {
		fmt.Printf("failed to load spec: %v\n", err)
		os.Exit(1)
	}

	swagger := doc.Spec()

	for path, pathItem := range swagger.Paths.Paths {
		fmt.Printf("Processing path: %s\n", path)
		for method, operation := range pathItem.Operations() {
			fmt.Printf("  Method: %s\n", method)

			for statusCode, response := range operation.Responses.StatusCodeResponses {
				fmt.Printf("    Status Code: %d\n", statusCode)
				schema := response.Schema
				if schema != nil {
					dereferencedSchema, err := dereferenceSchema(schema)
					if err != nil {
						fmt.Printf("      Failed to dereference schema: %v\n", err)
					} else {
						fmt.Printf("      Dereferenced Schema: %+v\n", dereferencedSchema)
					}
				} else {
					fmt.Println("      No schema found")
				}
			}
		}
	}
}
