package model

import "fmt"

type QueryParam struct {
	Query  string
	Author string
}

type QueryBuilder interface {
	BuildWhereClause() string
}

func (qp QueryParam) BuildWhereClause() string {
	if qp.Query == "" && qp.Author == "" {
		return ""
	}

	clause := "WHERE "
	if qp.Query != "" {
		clause += fmt.Sprintf(" title_body_vectors @@ TO_TSQUERY('%s') ", qp.Query)
	}

	if qp.Query != "" && qp.Author != "" {
		clause += "OR "
	}

	if qp.Author != "" {
		clause += fmt.Sprintf("author_vectors @@ TO_TSQUERY('%s') ", qp.Author)
	}

	return clause
}
