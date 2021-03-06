package updownapp

import (
	"errors"
	"github.com/couchbaselabs/go-couchbase"
	"strings"
	"time"
)

type IPresentation interface {
	Title() string
	SetTitle(string)
	CreatedAt() time.Time
	SetCreatedAt(time.Time)
	UpVotes() int
	SetUpVotes(int)
	DownVotes() int
	SetDownVotes(int)
	Persisted() bool
	SetPersisted(bool)
	Save() error
}

type Presentation struct {
	ID                string
	PresentationTitle string
	Created           time.Time
	Upvotes           int
	Downvotes         int
	IsPersisted       bool
}

func NewPresentation() IPresentation {
	return &Presentation{"", "", time.Now(), 0, 0, false}
}

func (p *Presentation) Key() string {
	return p.ID
}

func (p *Presentation) SetKey(key string) {
	p.ID = key
}

func (p *Presentation) Title() string {
	return p.PresentationTitle
}

func (p *Presentation) SetTitle(title string) {
	p.PresentationTitle = title
	p.ID = title
}

func (p *Presentation) CreatedAt() time.Time {
	return p.Created
}

func (p *Presentation) SetCreatedAt(createdAt time.Time) {
	p.Created = createdAt
}

func (p *Presentation) UpVotes() int {
	return p.Upvotes
}

func (p *Presentation) SetUpVotes(votes int) {
	p.Upvotes = votes
}

func (p *Presentation) DownVotes() int {
	return p.Downvotes
}

func (p *Presentation) SetDownVotes(downVotes int) {
	p.Downvotes = downVotes
}

func (p *Presentation) Persisted() bool {
	return p.IsPersisted
}

func (p *Presentation) SetPersisted(persisted bool) {
	p.IsPersisted = persisted
}

type PresentationIndex struct {
	Keys []string
}

func (p *Presentation) Save() error {
	factory := NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		return errors.New("Bucket should not be nil")
	}

	if !p.IsPersisted {
		indexKey := "Index"
		var index PresentationIndex
		err := bucket.Get(indexKey, &index)
		if err != nil {
			if strings.Contains(err.Error(), "Not found") {
				index = PresentationIndex{[]string{}}
			} else {
				return err
			}
		}

		keyExists := false
		for _, key := range index.Keys {
			if key == p.ID {
				keyExists = true
				break
			}
		}

		if !keyExists {
			index.Keys = append(index.Keys, p.ID)

			if err := bucket.Set(indexKey, 0, index); err != nil {
				return err
			}
		}

		p.IsPersisted = true
	}

	if err := bucket.Set(p.ID, 0, p); err != nil {
		return err
	}

	return nil
}

func FindPresentation(key string) (IPresentation, error) {
	bucket, err := getBucket()
	if err != nil {
		return nil, err
	}

	var presentation Presentation
	if err := bucket.Get(key, &presentation); err != nil {
		return nil, err
	}

	return &presentation, nil
}

func FindAllPresentations() ([]IPresentation, error) {
	indexKey := "Index"
	var index PresentationIndex
	bucket, err := getBucket()
	if err != nil {
		return nil, err
	}

	err = bucket.Get(indexKey, &index)
	if err != nil {
		if strings.Contains(err.Error(), "Not found") {
			index = PresentationIndex{[]string{}}
		} else {
			return nil, err
		}
	}

	var result []IPresentation
	for _, key := range index.Keys {
		presentation, err := FindPresentation(key)
		if err != nil {
			return nil, err
		}

		result = append(result, presentation)
	}

	return result, nil
}

func getBucket() (*couchbase.Bucket, error) {
	factory := NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		return nil, errors.New("Bucket should not be nil")
	}

	return bucket, nil
}
