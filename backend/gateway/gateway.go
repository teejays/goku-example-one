package gateway

import (
	"fmt"
	"io/ioutil"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	resolver "github.com/teejays/goku/example/backend/goku.generated/graphql"
)

// type query struct{}

// func (_ *query) Hello(ctx context.Context) (string, error) {
// 	return "Hello, world!", nil
// }

// type HelloArgs struct {
// 	Name string
// }

// func (_ *query) HelloWithArgs(ctx context.Context, args HelloArgs) (string, error) {
// 	return fmt.Sprintf("Hello, %s!", args.Name), nil
// }

// func main() {
// 	// Hacky path
// 	schemaFilePath := "example/gateway/goku.generated/graphql/schema.generated.graphql"
// 	err := StartServer("127.0.0.1", 8081, schemaFilePath)
// 	panics.IfError(err, "Starting GraphQL Server")
// }

func StartServer(addr string, port int, schemaFilePath string) error {
	if port < 1 {
		port = 8080
	}

	// Sample Schema
	/*
		s := `
				type Query {
					hello: String!
					helloWithArgs(name: String!): String!
				}
				type HelloArgs {
					Name: String!
				}
		`
	*/

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema, err := ParseSchemaFile(schemaFilePath, &resolver.Resolver{}, opts...)
	if err != nil {
		return err
	}
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	return http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
}

// TODO(): Contribute this
func ParseSchemaFile(schemaFilePath string, resolver interface{}, opts ...graphql.SchemaOpt) (*graphql.Schema, error) {
	schemaData, err := ioutil.ReadFile(schemaFilePath)
	if err != nil {
		return nil, err
	}
	schemaString := string(schemaData)
	schema, err := graphql.ParseSchema(schemaString, resolver, opts...)
	if err != nil {
		return nil, err
	}
	return schema, nil

}
