package inmem

import (
	"github.com/hashicorp/go-memdb"
)

func InitSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"events": {
				Name: "events",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:   "id",
						Unique: true,
						Indexer: &memdb.StringFieldIndex{
							Field: "ID",
						},
					},
					"source": {
						Name:   "source",
						Unique: false,
						Indexer: &memdb.StringFieldIndex{
							Field: "Source",
						},
					},
				},
			},
		},
	}
}
