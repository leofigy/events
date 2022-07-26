package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


import (
	"github.com/hashicorp/go-memdb"
	"github.com/leofigy/events/graph/inmem"
)

type Resolver struct{
	cfg ServerConfig
	memDb *memdb.MemDB
}

type ServerConfig struct {
	MemMode bool
}

func NewResolver( cfg ServerConfig) (rs *Resolver, err error) {

	var db *memdb.MemDB

	if cfg.MemMode {
		db, err = memdb.NewMemDB(inmem.InitSchema())
		if err != nil {
			return nil, err
		}
	}

	rs = &Resolver{
		cfg: cfg,
		memDb: db,
	}

	return
}
