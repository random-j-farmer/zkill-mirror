package db

import (
	"testing"

	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/zkb"
)

func TestStoreAndLookup(t *testing.T) {
	InitDB("testdb.x")
	defer DB.Close()

	km := &zkb.Killmail{
		KillID: 666,
	}
	err := IndexKillmail(bobstore.Ref{Fno: 666, Pos: 0xDEADBEEF}, km)
	if err != nil {
		t.Error("StoreKillmail", err)
	}

	/*
		m2, missing, err := LookupIDs([]string{"random j farmer", "rixx javix", "skir skor", "mynxee"})
		if err != nil {
			t.Error("LookupIDs", err)
		}

		if !reflect.DeepEqual([]string{"mynxee"}, missing) {
			t.Errorf("LookupIDs: expected: mynxee but was: %#v", missing)
		}

		if !reflect.DeepEqual(m, m2) {
			t.Errorf("LookupIDs: expected: %#v but was: %v", m, m2)
		}
	*/
}
