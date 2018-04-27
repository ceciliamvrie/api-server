package postgres_test

import (
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist"
	"github.com/techmexdev/lineuplist/pkg/postgres"
)

func TestFestivalSave(t *testing.T) {
	_, _, migrateDown := storeFestivals(t)
	defer migrateDown()
}

func TestLoadAll(t *testing.T) {
	ff, fStore, migrateDown := storeFestivals(t)
	defer migrateDown()

	storedFf, err := fStore.LoadAll()
	if err != nil {
		t.Fatalf("error loading all festivals: %s", err)
	}

	for i := range storedFf {
		log.Printf("%#v", storedFf[i])
		if storedFf[i].Name != ff[i].Name || !lineupEqual(storedFf[i].Lineup, ff[i].Lineup) {
			log.Fatalf("error comparing stored festivals: have %v, want %v", storedFf[i], ff[i])
		}
	}
}

func TestLoad(t *testing.T) {
	ff, fStore, migrateDown := storeFestivals(t)
	defer migrateDown()

	acl, err := fStore.Load("Austin City Limits")
	if err != nil {
		t.Fatalf("error loading festival: %s", err)
	}

	if !lineupEqual(acl.Lineup, ff[0].Lineup) {
		t.Fatalf("error loading festival: have %v, want %v", acl.Lineup, ff[0].Lineup)
	}
}

func TestFromArtist(t *testing.T) {
	_, fStore, migrateDown := storeFestivals(t)
	defer migrateDown()

	gorFf, err := fStore.FromArtist("Gorillaz")
	if err != nil {
		t.Fatal(err)
	}

	var containsACL, containsLev bool
	for _, f := range gorFf {
		if f.Name == "Austin City Limits" {
			containsACL = true
		}
		if f.Name == "Levitation" {
			containsLev = true
		}
	}
	if !containsACL || !containsLev {
		t.Fatalf("expecting %v to contain festivals: 'Austin City Limits', and 'Levitation'", gorFf)
	}
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

	return ff, fStore, func() {
		postgres.MigrateDown("file://../../migrations", os.Getenv("PG_TEST_DSN"))
	}
}

func lineupEqual(a, b []lineuplist.ArtistPreview) bool {
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
