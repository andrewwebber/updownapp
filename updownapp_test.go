package updownapp_test

import (
	"github.com/andrewwebber/updownapp"
	"log"
	"testing"
)

func TestGetConnection(t *testing.T) {
	factory := updownapp.NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		t.Fatal("Bucket should not be nil")
	}
}

func TestSave(t *testing.T) {
	Save(t, "TestSave")
}

func TestFind(t *testing.T) {
	key := "TestFind"
	Save(t, key)
	presentation, err := updownapp.FindPresentation(key)
	if err != nil {
		t.Fatal(err)
	}

	if presentation == nil {
		t.Fatal("Expected a presentation")
	}

	log.Println(presentation)
}

func TestFindAll(t *testing.T) {
	key := "TestFindAll"
	Save(t, key)
	presentations, err := updownapp.FindAllPresentations()
	if err != nil {
		t.Fatal(err)
	}

	if presentations == nil {
		t.Fatal("Expected to find presentations")
	}

	for _, presentation := range presentations {
		log.Println(presentation)
	}
}

type FatalLogger interface {
	Fatal(...interface{})
}

func Save(t FatalLogger, title string) {
	factory := updownapp.NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		t.Fatal("Bucket should not be nil")
	}

	presentation := updownapp.NewPresentation()
	presentation.SetTitle(title)
	presentation.Save()
}
