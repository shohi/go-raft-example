package server

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/hashicorp/raft"
	"github.com/shohi/go-raft-example/pkg/store"
)

type command struct {
	Op    string `json:"op,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type raftFSM struct {
	mu    sync.Mutex
	raft  *raftNode
	store store.Store
}

// Apply applies a Raft log entry to the key-value store.
func (f *raftFSM) Apply(l *raft.Log) interface{} {
	var c command
	if err := json.Unmarshal(l.Data, &c); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}

	switch c.Op {
	case "set":
		return f.applySet(c.Key, c.Value)
	case "delete":
		return f.applyDelete(c.Key)
	default:
		panic(fmt.Sprintf("unrecognized command op: %s", c.Op))
	}
}

func (f *raftFSM) applySet(key, value string) interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.store.Set(key, value)

	return nil
}

func (f *raftFSM) applyDelete(key string) interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.store.Delete(key)

	return nil
}

type fsmSnapshot struct {
	store map[string]string
}

// Snapshot returns a snapshot of the key-value store.
func (f *raftFSM) Snapshot() (raft.FSMSnapshot, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Clone the map.
	o := make(map[string]string, f.store.Len())
	f.store.ForEach(func(k, v string) error {
		o[k] = v
		return nil
	})

	return &fsmSnapshot{store: o}, nil
}

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := json.Marshal(f.store)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
	}

	return err
}

func (f *fsmSnapshot) Release() {}
