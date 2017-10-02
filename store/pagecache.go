package store

import (
	"io"
	"time"
)

/*
 *Store is used to abstract the underlying database
 *(bbolt). Pagecache is  used to cache metadata from
 *sites recently visited by the crawler. After duration specified
 *in global project config, the cache is dumped and unique urls
 *are persisted to the database, along with the metadata.
 */

type Store interface {
	Begin(rw io.ReadWriter, timeout time.Duration) error
	NewBucket()
}

type BoltStore struct{}

func (b *BoltStore) Open()
