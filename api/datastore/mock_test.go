package datastore

import (
	"github.com/IronFunctions2/functions/api/datastore/internal/datastoretest"
	"testing"
)

func TestDatastore(t *testing.T) {
	datastoretest.Test(t, NewMock())
}
