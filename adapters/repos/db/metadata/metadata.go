package metadata

import (
	"fmt"

	"github.com/gogo/protobuf/test/indeximport-issue72/index"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/flat"
)

type metadata struct {
	index flat
}

func (m *metadata) filter() error {
	return fmt.Errorf("Goon")
}