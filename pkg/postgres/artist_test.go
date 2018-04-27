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

func TestLoadAndLoadAll(t *testing.T) {
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

	gol, err := aStore.Load("Gorillaz")
	if err != nil {
		t.Fatalf("error loading artist: %s", err)
	}

	if gol.Name != "Gorillaz" {
		t.Fatalf("error loading artist: have %v, want %v", gol, aa[1])
	}
}

func TestArtistsFestivals(t *testing.T) {
	_, _, migrateDown := storeFestivals(t)
	defer migrateDown()

	aStore := postgres.NewArtistStorage(os.Getenv("PG_TEST_DSN"))

	gol, err := aStore.Load("Gorillaz")
	if err != nil {
		t.Fatalf("error loading artist: %s", err)
	}

	expFf := []lineuplist.FestivalPreview{{Name: "Austin City Limits"}, {Name: "Levitation"}}
	if gol.Festivals[0].Name != expFf[0].Name || gol.Festivals[1].Name != expFf[1].Name {
		t.Fatalf("error retrieving artist's festivals: have %v, want %v", gol.Festivals, expFf)
	}
}

func storeArtists(t *testing.T) (aa []lineuplist.Artist, as lineuplist.ArtistStorage, migrateDown func()) {
	aa = []lineuplist.Artist{
		{Name: "Kanye West"},
		{Name: "Gorillaz"},
		{Name: "Zoé"},
	}
	postgres.MigrateUp("file://../../migrations", os.Getenv("PG_TEST_DSN"))

	aStore := postgres.NewArtistStorage(os.Getenv("PG_TEST_DSN"))

	for _, a := range aa {
		_, err := aStore.Save(a)
		if err != nil {
			t.Fatalf("error saving artist %v: %s", a, err)
		}
	}

	return aa, aStore, func() { postgres.MigrateDown("file://../../migrations", os.Getenv("PG_TEST_DSN")) }
}

func storeFestivals(t *testing.T) (ff []lineuplist.Festival, fs lineuplist.FestivalStorage, migrateDown func()) {
	ff = []lineuplist.Festival{
		{
			Name:      "Austin City Limits",
			StartDate: time.Now(), EndDate: time.Now(),
			Country: "United States", State: "Tx", City: "Austin",
			Lineup: []lineuplist.ArtistPreview{{Name: "Red Hot Chilli Peppers"}, {Name: "Gorillaz"}, {Name: "Jay-Z"}},
		},
		{
			Name:      "Levitation",
			StartDate: time.Now(), EndDate: time.Now(),
			Country: "United States", State: "Tx", City: "Austin",
			Lineup: []lineuplist.ArtistPreview{{Name: "Gorillaz"}, {Name: "The Octopus Project"}, {Name: "Ariel Pink"}},
		},
	}

	postgres.MigrateUp("file://../../migrations", os.Getenv("PG_TEST_DSN"))

	fStore := postgres.NewFestivalStorage(os.Getenv("PG_TEST_DSN"))

	for _, f := range ff {
		_, err := fStore.Save(f)
		if err != nil {
			t.Fatal(err)
		}
	}

	return ff, fStore,
		func() {
			postgres.MigrateDown("file://../../migrations", os.Getenv("PG_TEST_DSN"))
		}
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
