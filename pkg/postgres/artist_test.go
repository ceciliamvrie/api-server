package postgres_test

import (
	"os"
	"testing"
	"time"

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

func TestFromFestival(t *testing.T) {
	ff, _, migrateDown := storeFestivals(t)
	defer migrateDown()

	aStore := postgres.NewArtistStorage(os.Getenv("PG_TEST_DSN"))

	aclAa, err := aStore.FromFestival("Austin City Limits")
	if err != nil {
		t.Fatal(err)
	}

	if !lineupEqual(ff[0].Lineup, aclAa) {
		t.Fatalf("error retriving artists from 'Austin City Limits': have %v, want %v", aclAa, ff[0].Lineup)
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

func storeFestivals(t *testing.T) (ff []lineuplist.Festival, fs lineuplist.FestivalStorage, migrateDown func()) {
	ff = []lineuplist.Festival{
		{
			Name:      "Austin City Limits",
			StartDate: time.Now(), EndDate: time.Now(),
			Country: "United States", State: "Tx", City: "Austin",
			Lineup: []lineuplist.Artist{{Name: "Red Hot Chilli Peppers"}, {Name: "Gorillaz"}, {Name: "Jay-Z"}},
		},
		{
			Name:      "Levitation",
			StartDate: time.Now(), EndDate: time.Now(),
			Country: "United States", State: "Tx", City: "Austin",
			Lineup: []lineuplist.Artist{{Name: "Gorillaz"}, {Name: "The Octopus Project"}, {Name: "Ariel Pink"}},
		},
	}

	postgres.MigrateUp(os.Getenv("PG_TEST_DSN"))

	fStore := postgres.NewFestivalStorage(os.Getenv("PG_TEST_DSN"))

	for _, f := range ff {
		_, err := fStore.Save(f)
		if err != nil {
			t.Fatal(err)
		}
	}

	return ff, fStore, func() { postgres.MigrateDown(os.Getenv("PG_TEST_DSN")) }
}

func lineupEqual(a, b []lineuplist.Artist) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(a); j++ {
			if a[i].Name == b[i].Name {
				found = true
				continue
			}
		}
		if !found {
			return false
		}
	}

	return true
}
