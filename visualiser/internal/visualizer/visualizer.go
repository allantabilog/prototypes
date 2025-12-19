package visualizer

import (
	"encoding/json"
	"time"
)

// Visualizable represents any data structure that can be visualized
type Visualizable interface {
	// Snapshot returns the current state of the data structure as JSON
	Snapshot() (*Snapshot, error)
	// GetID returns a unique identifier for this data structure instance
	GetID() string
	// GetType returns the type of data structure (e.g., "list", "stack", "queue", "tree")
	GetType() string
}

// Snapshot represents a point-in-time state of a data structure
type Snapshot struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// Operation represents an operation performed on a data structure
type Operation struct {
	ID            string                 `json:"id"`
	DataStructure string                 `json:"dataStructure"`
	Type          string                 `json:"type"` // "insert", "delete", "search", etc.
	Parameters    map[string]interface{} `json:"parameters"`
	Timestamp     time.Time              `json:"timestamp"`
	Before        *Snapshot              `json:"before,omitempty"`
	After         *Snapshot              `json:"after,omitempty"`
}

// Visualizer manages multiple data structures and their visualizations
type Visualizer struct {
	dataStructures map[string]Visualizable
	operations     []Operation
	subscribers    []chan Operation
}

// NewVisualizer creates a new visualizer instance
func NewVisualizer() *Visualizer {
	return &Visualizer{
		dataStructures: make(map[string]Visualizable),
		operations:     make([]Operation, 0),
		subscribers:    make([]chan Operation, 0),
	}
}

// Register adds a data structure to the visualizer
func (v *Visualizer) Register(ds Visualizable) {
	v.dataStructures[ds.GetID()] = ds
}

// Unregister removes a data structure from the visualizer
func (v *Visualizer) Unregister(id string) {
	delete(v.dataStructures, id)
}

// RecordOperation records an operation and notifies subscribers
func (v *Visualizer) RecordOperation(op Operation) {
	v.operations = append(v.operations, op)
	
	// Notify all subscribers
	for _, ch := range v.subscribers {
		select {
		case ch <- op:
		default:
			// Channel is full, skip this subscriber
		}
	}
}

// Subscribe returns a channel that receives operation updates
func (v *Visualizer) Subscribe() <-chan Operation {
	ch := make(chan Operation, 100) // Buffered channel
	v.subscribers = append(v.subscribers, ch)
	return ch
}

// GetSnapshot returns the current snapshot of a specific data structure
func (v *Visualizer) GetSnapshot(id string) (*Snapshot, error) {
	ds, exists := v.dataStructures[id]
	if !exists {
		return nil, ErrDataStructureNotFound
	}
	return ds.Snapshot()
}

// GetAllSnapshots returns snapshots of all registered data structures
func (v *Visualizer) GetAllSnapshots() (map[string]*Snapshot, error) {
	snapshots := make(map[string]*Snapshot)
	
	for id, ds := range v.dataStructures {
		snapshot, err := ds.Snapshot()
		if err != nil {
			return nil, err
		}
		snapshots[id] = snapshot
	}
	
	return snapshots, nil
}

// GetOperations returns all recorded operations
func (v *Visualizer) GetOperations() []Operation {
	return v.operations
}

// GetOperationsForDataStructure returns operations for a specific data structure
func (v *Visualizer) GetOperationsForDataStructure(id string) []Operation {
	var ops []Operation
	for _, op := range v.operations {
		if op.DataStructure == id {
			ops = append(ops, op)
		}
	}
	return ops
}

// ToJSON converts a snapshot to JSON bytes
func (s *Snapshot) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// ToJSON converts an operation to JSON bytes
func (o *Operation) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}