package internal

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	rpc "github.com/hash167/yt-dl-rpc/api/rpc"
	"github.com/hash167/yt-dl-rpc/server/config"
)

// In-Memory Thread-Safe Key-Value Storage with optional persistence
type MemoryDB struct {
	table sync.Map
}

type Session struct {
	Processes []*rpc.ProcessResponse
}

// Get a process pointer given its id
func (m *MemoryDB) Get(id string) (*Process, error) {
	entry, ok := m.table.Load(id)
	if !ok {
		return nil, errors.New("no process found for the given key")
	}
	return entry.(*Process), nil
}

// Store a pointer of a process and return its id
func (m *MemoryDB) Set(process *Process) string {
	id := uuid.NewString()
	m.table.Store(id, process)
	process.Id = id
	return id
}

// Update a process info/metadata, given the process id
func (m *MemoryDB) UpdateInfo(id string, info *rpc.DownloadInfo) error {
	entry, ok := m.table.Load(id)
	if ok {
		entry.(*Process).Info = info
		m.table.Store(id, entry)
		return nil
	}
	return fmt.Errorf("can't update row with id %s", id)
}

// Update a process progress data, given the process id
// Used for updating completition percentage or ETA.
//
// Deprecated: will be removed anytime soon.
func (m *MemoryDB) UpdateProgress(id string, progress *rpc.DownloadProgress) error {
	entry, ok := m.table.Load(id)
	if ok {
		entry.(*Process).Progress = progress
		m.table.Store(id, entry)
		return nil
	}
	return fmt.Errorf("can't update row with id %s", id)
}

// Removes a process progress, given the process id
func (m *MemoryDB) Delete(id string) {
	m.table.Delete(id)
}

func (m *MemoryDB) Keys() *[]string {
	running := []string{}
	m.table.Range(func(key, value any) bool {
		running = append(running, key.(string))
		return true
	})
	return &running
}

// Returns a slice of all currently stored processes progess
func (m *MemoryDB) All() []*rpc.ProcessResponse {
	running := []*rpc.ProcessResponse{}
	m.table.Range(func(key, value any) bool {
		running = append(running, &rpc.ProcessResponse{
			Id:       key.(string),
			Info:     value.(*Process).Info,
			Progress: value.(*Process).Progress,
			Output:   value.(*Process).Output,
			Params:   value.(*Process).Params,
		})
		return true
	})
	return running
}

// WIP: Persist the database in a single file named "session.dat"
func (m *MemoryDB) Persist() error {
	running := m.All()

	sf := filepath.Join(config.Instance().SessionFilePath, "session.dat")

	fd, err := os.Create(sf)
	if err != nil {
		return errors.Join(errors.New("failed to persist session"), err)
	}

	session := Session{
		Processes: running,
	}

	err = gob.NewEncoder(fd).Encode(session)
	if err != nil {
		return errors.Join(errors.New("failed to persist session"), err)
	}

	return nil
}

// WIP: Restore a persisted state
func (m *MemoryDB) Restore(logger *slog.Logger) {
	fd, err := os.Open("session.dat")
	if err != nil {
		return
	}

	session := Session{}

	err = gob.NewDecoder(fd).Decode(&session)
	if err != nil {
		return
	}

	for _, proc := range session.Processes {
		restored := &Process{
			Id:       proc.Id,
			Url:      proc.Info.URL,
			Info:     proc.Info,
			Progress: proc.Progress,
			Output:   proc.Output,
			Params:   proc.Params,
			Logger:   logger,
		}

		m.table.Store(proc.Id, restored)

		if restored.Progress.Percentage != "-1" {
			go restored.Start()
		}
	}
}
