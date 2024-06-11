package server

import (
	"fmt"
	"sync"
)

type Log struct {
	// A mutex locks whatever thread this resource is on to prevent other resources from accessing it?
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// appends to end of log
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	// I think the purpose of defer here is to unlock this resource when the function is done running
	defer c.mu.Unlock()
	// this should put us at the end of the log so we can just append updates
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
