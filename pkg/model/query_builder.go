package model

type QueryBuilder interface {
	BuildWhereClause() string
}
