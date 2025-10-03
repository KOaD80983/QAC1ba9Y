// 代码生成时间: 2025-10-04 02:30:22
// graphql_api.go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/graphql-go/graphql"
)

// 初始化GraphQL Schema
func init() {
    fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return "world", nil
            },
        },
    }
    schemaConfig := graphql.SchemaConfig{
        Query: graphql.ObjectConfig{Name: "RootQuery", Fields: fields},
    }
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("Failed to create new GraphQL schema, error: %v", err)
    }
    graphql.MustParseSchema(schema)
}

// GraphQLHandler 处理GraphQL请求
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    query := r.Form.Get("query")
    if query == "" {
        http.Error(w, "Must provide query string", http.StatusBadRequest)
        return
    }

    result := graphql.Do(graphql.Params{
        Schema:        schema,
        RequestString: query,
    })
    if len(result.Errors) > 0 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, result.Errors)
    } else {
        fmt.Fprintln(w, result.Data)
    }
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.URLFormat)

    r.Post("/graphql", GraphQLHandler)

    log.Printf("Server started on port 3000")
    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Fatal(err)
    }
}
