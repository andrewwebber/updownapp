package updownapp

import (
	"github.com/couchbaselabs/go-couchbase"
	"log"
)

type CouchbaseConnectionFactory struct {
	bucket *couchbase.Bucket
}

func NewCouchbaseConnectionFactory() *CouchbaseConnectionFactory {
	c, err := couchbase.Connect("http://localhost:8091/")
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	pool, err := c.GetPool("default")
	if err != nil {
		log.Fatalf("Error getting pool:  %v", err)
	}

	bucket, err := pool.GetBucket("default")
	if err != nil {
		log.Fatalf("Error getting bucket:  %v", err)
	}

	return &CouchbaseConnectionFactory{bucket}
}

func (f *CouchbaseConnectionFactory) GetDefaultConnection() *couchbase.Bucket {
	return f.bucket
}
