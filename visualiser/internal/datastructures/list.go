package datastructures

import (
	"fmt"
	"time"

	"github.com/allantabilog/visualiser/internal/visualizer"
)

// VisualizableList is a dynamic array that supports visualization
type VisualizableList struct {
	id         string
	items      []interface{}
	visualizer *visualizer.Visualizer
}

// NewVisualizableList creates a new visualizable list
func NewVisualizableList(id string, viz *visualizer.Visualizer) *VisualizableList {
	list := &VisualizableList{
		id:         id,
		items:      make([]interface{}, 0),
		visualizer: viz,
	}
	
	if viz != nil {
		viz.Register(list)
	}
	
	return list
}

// GetID returns the unique identifier for this list
func (l *VisualizableList) GetID() string {
	return l.id
}

// GetType returns the type of data structure
func (l *VisualizableList) GetType() string {
	return "list"
}

// Snapshot returns the current state of the list
func (l *VisualizableList) Snapshot() (*visualizer.Snapshot, error) {
	data := map[string]interface{}{
		"items":  l.items,
		"length": len(l.items),
	}
	
	metadata := map[string]interface{}{
		"capacity": cap(l.items),
	}
	
	return &visualizer.Snapshot{
		ID:        l.id,
		Type:      l.GetType(),
		Timestamp: time.Now(),
		Data:      data,
		Metadata:  metadata,
	}, nil
}

// recordOperation records an operation with before and after snapshots
func (l *VisualizableList) recordOperation(opType string, params map[string]interface{}) {
	if l.visualizer == nil {
		return
	}
	
	before, _ := l.Snapshot()
	
	// Perform the actual operation here (this will be called after the operation)
	// So we need to get the after snapshot in the calling method
	
	operation := visualizer.Operation{
		ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
		DataStructure: l.id,
		Type:          opType,
		Parameters:    params,
		Timestamp:     time.Now(),
		Before:        before,
		// After will be set by the calling method
	}
	
	l.visualizer.RecordOperation(operation)
}

// Append adds an element to the end of the list
func (l *VisualizableList) Append(item interface{}) {
	before, _ := l.Snapshot()
	
	l.items = append(l.items, item)
	
	if l.visualizer != nil {
		after, _ := l.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
			DataStructure: l.id,
			Type:          "append",
			Parameters:    map[string]interface{}{"item": item},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		l.visualizer.RecordOperation(operation)
	}
}

// Insert adds an element at a specific index
func (l *VisualizableList) Insert(index int, item interface{}) error {
	if index < 0 || index > len(l.items) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	
	before, _ := l.Snapshot()
	
	// Insert the item
	l.items = append(l.items[:index], append([]interface{}{item}, l.items[index:]...)...)
	
	if l.visualizer != nil {
		after, _ := l.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
			DataStructure: l.id,
			Type:          "insert",
			Parameters:    map[string]interface{}{"index": index, "item": item},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		l.visualizer.RecordOperation(operation)
	}
	
	return nil
}

// Delete removes an element at a specific index
func (l *VisualizableList) Delete(index int) (interface{}, error) {
	if index < 0 || index >= len(l.items) {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}
	
	before, _ := l.Snapshot()
	
	item := l.items[index]
	l.items = append(l.items[:index], l.items[index+1:]...)
	
	if l.visualizer != nil {
		after, _ := l.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
			DataStructure: l.id,
			Type:          "delete",
			Parameters:    map[string]interface{}{"index": index},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		l.visualizer.RecordOperation(operation)
	}
	
	return item, nil
}

// Get returns the element at a specific index
func (l *VisualizableList) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(l.items) {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}
	
	// Record search operation
	if l.visualizer != nil {
		snapshot, _ := l.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
			DataStructure: l.id,
			Type:          "get",
			Parameters:    map[string]interface{}{"index": index},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		l.visualizer.RecordOperation(operation)
	}
	
	return l.items[index], nil
}

// Length returns the number of elements in the list
func (l *VisualizableList) Length() int {
	return len(l.items)
}

// IsEmpty returns true if the list is empty
func (l *VisualizableList) IsEmpty() bool {
	return len(l.items) == 0
}

// Clear removes all elements from the list
func (l *VisualizableList) Clear() {
	before, _ := l.Snapshot()
	
	l.items = l.items[:0] // Keep the underlying array
	
	if l.visualizer != nil {
		after, _ := l.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", l.id, time.Now().UnixNano()),
			DataStructure: l.id,
			Type:          "clear",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		l.visualizer.RecordOperation(operation)
	}
}