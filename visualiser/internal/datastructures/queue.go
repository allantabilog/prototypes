package datastructures

import (
	"fmt"
	"time"

	"github.com/allantabilog/visualiser/internal/visualizer"
)

// VisualizableQueue implements a FIFO queue with visualization support
type VisualizableQueue struct {
	id         string
	items      []interface{}
	visualizer *visualizer.Visualizer
}

// NewVisualizableQueue creates a new visualizable queue
func NewVisualizableQueue(id string, viz *visualizer.Visualizer) *VisualizableQueue {
	queue := &VisualizableQueue{
		id:         id,
		items:      make([]interface{}, 0),
		visualizer: viz,
	}
	
	if viz != nil {
		viz.Register(queue)
	}
	
	return queue
}

// GetID returns the unique identifier for this queue
func (q *VisualizableQueue) GetID() string {
	return q.id
}

// GetType returns the type of data structure
func (q *VisualizableQueue) GetType() string {
	return "queue"
}

// Snapshot returns the current state of the queue
func (q *VisualizableQueue) Snapshot() (*visualizer.Snapshot, error) {
	data := map[string]interface{}{
		"items": q.items,
		"size":  len(q.items),
	}
	
	metadata := map[string]interface{}{
		"capacity": cap(q.items),
	}
	
	if len(q.items) > 0 {
		data["front"] = q.items[0]
		data["rear"] = q.items[len(q.items)-1]
	}
	
	return &visualizer.Snapshot{
		ID:        q.id,
		Type:      q.GetType(),
		Timestamp: time.Now(),
		Data:      data,
		Metadata:  metadata,
	}, nil
}

// Enqueue adds an element to the rear of the queue
func (q *VisualizableQueue) Enqueue(item interface{}) {
	before, _ := q.Snapshot()
	
	q.items = append(q.items, item)
	
	if q.visualizer != nil {
		after, _ := q.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", q.id, time.Now().UnixNano()),
			DataStructure: q.id,
			Type:          "enqueue",
			Parameters:    map[string]interface{}{"item": item},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		q.visualizer.RecordOperation(operation)
	}
}

// Dequeue removes and returns the front element from the queue
func (q *VisualizableQueue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	
	before, _ := q.Snapshot()
	
	item := q.items[0]
	q.items = q.items[1:]
	
	if q.visualizer != nil {
		after, _ := q.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", q.id, time.Now().UnixNano()),
			DataStructure: q.id,
			Type:          "dequeue",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		q.visualizer.RecordOperation(operation)
	}
	
	return item, nil
}

// Front returns the front element without removing it
func (q *VisualizableQueue) Front() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	
	// Record peek operation
	if q.visualizer != nil {
		snapshot, _ := q.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", q.id, time.Now().UnixNano()),
			DataStructure: q.id,
			Type:          "front",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		q.visualizer.RecordOperation(operation)
	}
	
	return q.items[0], nil
}

// Rear returns the rear element without removing it
func (q *VisualizableQueue) Rear() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	
	// Record peek operation
	if q.visualizer != nil {
		snapshot, _ := q.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", q.id, time.Now().UnixNano()),
			DataStructure: q.id,
			Type:          "rear",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		q.visualizer.RecordOperation(operation)
	}
	
	return q.items[len(q.items)-1], nil
}

// Size returns the number of elements in the queue
func (q *VisualizableQueue) Size() int {
	return len(q.items)
}

// IsEmpty returns true if the queue is empty
func (q *VisualizableQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Clear removes all elements from the queue
func (q *VisualizableQueue) Clear() {
	before, _ := q.Snapshot()
	
	q.items = q.items[:0]
	
	if q.visualizer != nil {
		after, _ := q.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", q.id, time.Now().UnixNano()),
			DataStructure: q.id,
			Type:          "clear",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		q.visualizer.RecordOperation(operation)
	}
}