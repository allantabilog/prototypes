package datastructures

import (
	"fmt"
	"time"

	"github.com/allantabilog/visualiser/internal/visualizer"
)

// VisualizableStack implements a LIFO stack with visualization support
type VisualizableStack struct {
	id         string
	items      []interface{}
	visualizer *visualizer.Visualizer
}

// NewVisualizableStack creates a new visualizable stack
func NewVisualizableStack(id string, viz *visualizer.Visualizer) *VisualizableStack {
	stack := &VisualizableStack{
		id:         id,
		items:      make([]interface{}, 0),
		visualizer: viz,
	}
	
	if viz != nil {
		viz.Register(stack)
	}
	
	return stack
}

// GetID returns the unique identifier for this stack
func (s *VisualizableStack) GetID() string {
	return s.id
}

// GetType returns the type of data structure
func (s *VisualizableStack) GetType() string {
	return "stack"
}

// Snapshot returns the current state of the stack
func (s *VisualizableStack) Snapshot() (*visualizer.Snapshot, error) {
	data := map[string]interface{}{
		"items": s.items,
		"size":  len(s.items),
	}
	
	metadata := map[string]interface{}{
		"capacity": cap(s.items),
	}
	
	if len(s.items) > 0 {
		data["top"] = s.items[len(s.items)-1]
	}
	
	return &visualizer.Snapshot{
		ID:        s.id,
		Type:      s.GetType(),
		Timestamp: time.Now(),
		Data:      data,
		Metadata:  metadata,
	}, nil
}

// Push adds an element to the top of the stack
func (s *VisualizableStack) Push(item interface{}) {
	before, _ := s.Snapshot()
	
	s.items = append(s.items, item)
	
	if s.visualizer != nil {
		after, _ := s.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", s.id, time.Now().UnixNano()),
			DataStructure: s.id,
			Type:          "push",
			Parameters:    map[string]interface{}{"item": item},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		s.visualizer.RecordOperation(operation)
	}
}

// Pop removes and returns the top element from the stack
func (s *VisualizableStack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	
	before, _ := s.Snapshot()
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	
	if s.visualizer != nil {
		after, _ := s.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", s.id, time.Now().UnixNano()),
			DataStructure: s.id,
			Type:          "pop",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		s.visualizer.RecordOperation(operation)
	}
	
	return item, nil
}

// Peek returns the top element without removing it
func (s *VisualizableStack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	
	// Record peek operation
	if s.visualizer != nil {
		snapshot, _ := s.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", s.id, time.Now().UnixNano()),
			DataStructure: s.id,
			Type:          "peek",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		s.visualizer.RecordOperation(operation)
	}
	
	return s.items[len(s.items)-1], nil
}

// Size returns the number of elements in the stack
func (s *VisualizableStack) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the stack is empty
func (s *VisualizableStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements from the stack
func (s *VisualizableStack) Clear() {
	before, _ := s.Snapshot()
	
	s.items = s.items[:0]
	
	if s.visualizer != nil {
		after, _ := s.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", s.id, time.Now().UnixNano()),
			DataStructure: s.id,
			Type:          "clear",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		s.visualizer.RecordOperation(operation)
	}
}