package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
)

type SimpleInputQueryResponse struct {
	User SimpleInputQueryUser `json:"user"`
}

type SimpleInputQueryUser struct {
	Id string `json:"id"`
}

func SimpleInputQuery(
	client graphql.Client,
	name string,
) (*SimpleInputQueryResponse, error) {
	variables := map[string]interface{}{
		"name": name,
	}

	var retval SimpleInputQueryResponse
	err := client.MakeRequest(
		nil,
		"SimpleInputQuery",
		`
query SimpleInputQuery ($name: String!) {
	user(query: {name:$name}) {
		id
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}