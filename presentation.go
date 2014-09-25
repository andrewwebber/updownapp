package updownapp

import (
	"errors"
	"time"
)

type IPresentation interface {
	Key() string
	SetKey(string)
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

func (p *Presentation) Save() error {
	factory := NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		return errors.New("Bucket should not be nil")
	}

	if err := bucket.Set(p.ID, 0, p); err != nil {
		return err
	}

	p.IsPersisted = true

	return nil
}

func FindPresentation(key string) (IPresentation, error) {
	factory := NewCouchbaseConnectionFactory()
	bucket := factory.GetDefaultConnection()
	if bucket == nil {
		return nil, errors.New("Bucket should not be nil")
	}

	var presentation Presentation
	if err := bucket.Get(key, &presentation); err != nil {
		return nil, err
	}

	return &presentation, nil
}

func FindAllPresentations() ([]IPresentation, error) {
	return nil, nil
}
