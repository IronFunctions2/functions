package bolt

import (
	"github.com/IronFunctions2/functions/api/datastore/internal/datastoretest"
	"net/url"
	"os"
	"testing"
)

const tmpBolt = "/tmp/bolt_fn_test.db"

func TestDatastore(t *testing.T) {
	u, err := url.Parse("bolt://" + tmpBolt)
	if err != nil {
		t.Fatalf("failed to parse url: %s", err)
	}
	ds, err := New(u)
	if err != nil {
		t.Fatalf("failed to create bolt datastore: %s", err)
	}
	datastoretest.Test(t, ds)
	os.Remove(tmpBolt)
}
