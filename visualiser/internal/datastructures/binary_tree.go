package datastructures

import (
	"fmt"
	"time"

	"github.com/allantabilog/visualiser/internal/visualizer"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value  interface{} `json:"value"`
	Left   *TreeNode   `json:"left,omitempty"`
	Right  *TreeNode   `json:"right,omitempty"`
	Parent *TreeNode   `json:"-"` // Don't include parent in JSON to avoid circular references
}

// VisualizableBinaryTree implements a binary tree with visualization support
type VisualizableBinaryTree struct {
	id         string
	root       *TreeNode
	size       int
	visualizer *visualizer.Visualizer
}

// NewVisualizableBinaryTree creates a new visualizable binary tree
func NewVisualizableBinaryTree(id string, viz *visualizer.Visualizer) *VisualizableBinaryTree {
	tree := &VisualizableBinaryTree{
		id:         id,
		root:       nil,
		size:       0,
		visualizer: viz,
	}
	
	if viz != nil {
		viz.Register(tree)
	}
	
	return tree
}

// GetID returns the unique identifier for this tree
func (t *VisualizableBinaryTree) GetID() string {
	return t.id
}

// GetType returns the type of data structure
func (t *VisualizableBinaryTree) GetType() string {
	return "binary_tree"
}

// nodeToMap converts a tree node to a map representation for JSON serialization
func (t *VisualizableBinaryTree) nodeToMap(node *TreeNode) map[string]interface{} {
	if node == nil {
		return nil
	}
	
	nodeMap := map[string]interface{}{
		"value": node.Value,
	}
	
	if node.Left != nil {
		nodeMap["left"] = t.nodeToMap(node.Left)
	}
	
	if node.Right != nil {
		nodeMap["right"] = t.nodeToMap(node.Right)
	}
	
	return nodeMap
}

// Snapshot returns the current state of the tree
func (t *VisualizableBinaryTree) Snapshot() (*visualizer.Snapshot, error) {
	data := map[string]interface{}{
		"size": t.size,
	}
	
	if t.root != nil {
		data["root"] = t.nodeToMap(t.root)
	}
	
	metadata := map[string]interface{}{
		"height": t.calculateHeight(t.root),
	}
	
	return &visualizer.Snapshot{
		ID:        t.id,
		Type:      t.GetType(),
		Timestamp: time.Now(),
		Data:      data,
		Metadata:  metadata,
	}, nil
}

// calculateHeight calculates the height of a tree rooted at the given node
func (t *VisualizableBinaryTree) calculateHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	
	leftHeight := t.calculateHeight(node.Left)
	rightHeight := t.calculateHeight(node.Right)
	
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Insert adds a value to the tree (as a binary search tree)
func (t *VisualizableBinaryTree) Insert(value interface{}) {
	before, _ := t.Snapshot()
	
	t.root = t.insertNode(t.root, nil, value)
	t.size++
	
	if t.visualizer != nil {
		after, _ := t.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", t.id, time.Now().UnixNano()),
			DataStructure: t.id,
			Type:          "insert",
			Parameters:    map[string]interface{}{"value": value},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		t.visualizer.RecordOperation(operation)
	}
}

// insertNode recursively inserts a node into the tree
func (t *VisualizableBinaryTree) insertNode(node *TreeNode, parent *TreeNode, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{
			Value:  value,
			Parent: parent,
		}
	}
	
	// For demonstration, we'll compare as strings
	// In a real implementation, you'd want proper comparison logic
	nodeStr := fmt.Sprintf("%v", node.Value)
	valueStr := fmt.Sprintf("%v", value)
	
	if valueStr < nodeStr {
		node.Left = t.insertNode(node.Left, node, value)
	} else {
		node.Right = t.insertNode(node.Right, node, value)
	}
	
	return node
}

// Search looks for a value in the tree
func (t *VisualizableBinaryTree) Search(value interface{}) bool {
	if t.visualizer != nil {
		snapshot, _ := t.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", t.id, time.Now().UnixNano()),
			DataStructure: t.id,
			Type:          "search",
			Parameters:    map[string]interface{}{"value": value},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		t.visualizer.RecordOperation(operation)
	}
	
	return t.searchNode(t.root, value)
}

// searchNode recursively searches for a value in the tree
func (t *VisualizableBinaryTree) searchNode(node *TreeNode, value interface{}) bool {
	if node == nil {
		return false
	}
	
	nodeStr := fmt.Sprintf("%v", node.Value)
	valueStr := fmt.Sprintf("%v", value)
	
	if valueStr == nodeStr {
		return true
	} else if valueStr < nodeStr {
		return t.searchNode(node.Left, value)
	} else {
		return t.searchNode(node.Right, value)
	}
}

// InOrderTraversal returns values in in-order traversal
func (t *VisualizableBinaryTree) InOrderTraversal() []interface{} {
	var result []interface{}
	t.inOrderHelper(t.root, &result)
	
	if t.visualizer != nil {
		snapshot, _ := t.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", t.id, time.Now().UnixNano()),
			DataStructure: t.id,
			Type:          "traversal_inorder",
			Parameters:    map[string]interface{}{"result": result},
			Timestamp:     time.Now(),
			Before:        snapshot,
			After:         snapshot, // No change in state
		}
		t.visualizer.RecordOperation(operation)
	}
	
	return result
}

// inOrderHelper performs in-order traversal recursively
func (t *VisualizableBinaryTree) inOrderHelper(node *TreeNode, result *[]interface{}) {
	if node != nil {
		t.inOrderHelper(node.Left, result)
		*result = append(*result, node.Value)
		t.inOrderHelper(node.Right, result)
	}
}

// Size returns the number of nodes in the tree
func (t *VisualizableBinaryTree) Size() int {
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *VisualizableBinaryTree) IsEmpty() bool {
	return t.root == nil
}

// Clear removes all nodes from the tree
func (t *VisualizableBinaryTree) Clear() {
	before, _ := t.Snapshot()
	
	t.root = nil
	t.size = 0
	
	if t.visualizer != nil {
		after, _ := t.Snapshot()
		operation := visualizer.Operation{
			ID:            fmt.Sprintf("%s_%d", t.id, time.Now().UnixNano()),
			DataStructure: t.id,
			Type:          "clear",
			Parameters:    map[string]interface{}{},
			Timestamp:     time.Now(),
			Before:        before,
			After:         after,
		}
		t.visualizer.RecordOperation(operation)
	}
}