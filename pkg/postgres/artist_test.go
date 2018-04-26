package postgres_test

import (
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist"
	"github.com/techmexdev/lineuplist/pkg/postgres"
)

func TestAritstStore(t *testing.T) {
	_, _, migrateDown := storeArtists(t)
	defer migrateDown()
}

func TestLoadAll(t *testing.T) {
	aa, aStore, migrateDown := storeArtists(t)
	defer migrateDown()

	storedA, err := aStore.LoadAll()
	if err != nil {
		t.Fatalf("error loading all artists: %s", err)
	}

	for i := range storedA {
		if storedA[i].Name != aa[i].Name {
			t.Fatalf("error comparing stored artists: have %v, want %v", storedA[i], aa[i])
		}
	}
}

func TestLoad(t *testing.T) {
	aa, aStore, migrateDown := storeArtists(t)
	defer migrateDown()

	lana, err := aStore.Load("Lana del Rey")
	if err != nil {
		t.Fatalf("error loading artist: %s", err)
	}

	if lana.Name != "Lana del Rey" {
		t.Fatalf("error loading artist: have %v, want %v", lana, aa[1])
	}
}

func storeArtists(t *testing.T) (aa []lineuplist.Artist, as lineuplist.ArtistStorage, migrateDown func()) {
	aa = []lineuplist.Artist{
		{Name: "Kanye West"},
		{Name: "Lana del Rey"},
		{Name: "Zoé"},
	}
	postgres.MigrateUp(os.Getenv("PG_TEST_DSN"))

	aStore := postgres.NewArtistStorage(os.Getenv("PG_TEST_DSN"))

	for _, a := range aa {
		_, err := aStore.Save(a)
		if err != nil {
			t.Fatalf("error saving artist %v: %s", a, err)
		}
	}

	return aa, aStore, func() { postgres.MigrateDown(os.Getenv("PG_TEST_DSN")) }
}
