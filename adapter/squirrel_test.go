package adapter

import (
	"d3/orm/entity"
	"d3/orm/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	query *query.Query
	sql   string
	args  []interface{}
}

var metaStub = &entity.MetaInfo{
	Fields:      map[string]*entity.FieldInfo{
		"id": {
			DbAlias:  "id",
		},
	},
	TableName:   "test_table",
}

var testCases = []testCase{
	{
		query: query.NewQuery(metaStub).AndWhere("id = ?", 1).OrWhere("id = ?", 3).Limit(1),
		sql:   "SELECT test_table.id as \"test_table.id\" FROM test_table WHERE (id = $1 OR id = $2) LIMIT 1",
		args:  []interface{}{1, 3},
	},
	{
		query: query.NewQuery(metaStub).AndWhere("id = ?", 1).OrWhere("id = ?", 3).Limit(1).
		Union(query.NewQuery(metaStub).AndWhere("id = ?", 5), "UNION ALL"),
		sql:   "SELECT test_table.id as \"test_table.id\" FROM test_table WHERE (id = $1 OR id = $2) LIMIT 1 UNION ALL SELECT test_table.id as \"test_table.id\" FROM test_table WHERE id = $3",
		args:  []interface{}{1, 3, 5},
	},
}

func TestQueryToSquirrelSql(t *testing.T) {
	adapter := SquirrelAdapter{}

	for _, tCase := range testCases {
		sql, args, _ := adapter.ToSql(tCase.query)
		assert.Equal(t, tCase.sql, sql)
		assert.Equal(t, tCase.args, args)
	}
}
