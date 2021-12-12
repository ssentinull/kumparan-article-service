package query_param_test

import (
	"testing"

	"github.com/ssentinull/kumparan-article-service/pkg/utils/query_param"
	"github.com/stretchr/testify/assert"
)

func TestValidBuildWhereClause(t *testing.T) {
	qp := query_param.QueryParam{
		Author: "author",
		Query:  "query",
	}

	expectedWhereClause := "WHERE  title_body_vectors @@ TO_TSQUERY('query') OR author_vectors @@ TO_TSQUERY('author') "
	whereClause := qp.BuildWhereClause()

	assert.Equal(t, expectedWhereClause, whereClause)
}

func TestValidSingularBuildWhereClause(t *testing.T) {
	qp := query_param.QueryParam{Query: "query"}

	expectedWhereClause := "WHERE  title_body_vectors @@ TO_TSQUERY('query') "
	whereClause := qp.BuildWhereClause()

	assert.Equal(t, expectedWhereClause, whereClause)
}

func TestInvalidBuildWhereClause(t *testing.T) {
	qp := query_param.QueryParam{}
	whereClause := qp.BuildWhereClause()

	assert.Empty(t, whereClause)
}
