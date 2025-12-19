package examples

import (
	"time"

	"github.com/allantabilog/visualiser/internal/datastructures"
	"github.com/allantabilog/visualiser/internal/visualizer"
)

// BubbleSortDemo demonstrates bubble sort algorithm with visualization
func BubbleSortDemo(viz *visualizer.Visualizer) {
	// Create a visualizable list with some unsorted data
	list := datastructures.NewVisualizableList("bubble_sort_demo", viz)
	
	// Add some unsorted numbers
	numbers := []interface{}{64, 34, 25, 12, 22, 11, 90}
	for _, num := range numbers {
		list.Append(num)
		time.Sleep(500 * time.Millisecond) // Pause to show each addition
	}
	
	// Perform bubble sort with visualization
	n := list.Length()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Get adjacent elements
			elem1, _ := list.Get(j)
			elem2, _ := list.Get(j + 1)
			
			// Convert to integers for comparison
			val1 := elem1.(int)
			val2 := elem2.(int)
			
			if val1 > val2 {
				// Swap elements
				list.Delete(j + 1)
				list.Delete(j)
				list.Insert(j, elem2)
				list.Insert(j + 1, elem1)
				
				// Pause to show the swap
				time.Sleep(1 * time.Second)
			}
		}
	}
}

// StackOperationsDemo demonstrates stack operations
func StackOperationsDemo(viz *visualizer.Visualizer) {
	stack := datastructures.NewVisualizableStack("stack_demo", viz)
	
	// Push some elements
	elements := []interface{}{"A", "B", "C", "D", "E"}
	for _, elem := range elements {
		stack.Push(elem)
		time.Sleep(800 * time.Millisecond)
	}
	
	// Peek at the top
	stack.Peek()
	time.Sleep(500 * time.Millisecond)
	
	// Pop some elements
	for i := 0; i < 3; i++ {
		stack.Pop()
		time.Sleep(800 * time.Millisecond)
	}
	
	// Push more elements
	stack.Push("F")
	time.Sleep(800 * time.Millisecond)
	stack.Push("G")
	time.Sleep(800 * time.Millisecond)
	
	// Clear the stack
	stack.Clear()
}

// QueueOperationsDemo demonstrates queue operations
func QueueOperationsDemo(viz *visualizer.Visualizer) {
	queue := datastructures.NewVisualizableQueue("queue_demo", viz)
	
	// Enqueue some elements
	customers := []interface{}{"Customer 1", "Customer 2", "Customer 3", "Customer 4"}
	for _, customer := range customers {
		queue.Enqueue(customer)
		time.Sleep(800 * time.Millisecond)
	}
	
	// Check front and rear
	queue.Front()
	time.Sleep(500 * time.Millisecond)
	queue.Rear()
	time.Sleep(500 * time.Millisecond)
	
	// Process some customers (dequeue)
	for i := 0; i < 2; i++ {
		queue.Dequeue()
		time.Sleep(1 * time.Second)
	}
	
	// Add more customers
	queue.Enqueue("Customer 5")
	time.Sleep(800 * time.Millisecond)
	queue.Enqueue("Customer 6")
	time.Sleep(800 * time.Millisecond)
	
	// Process remaining customers
	for !queue.IsEmpty() {
		queue.Dequeue()
		time.Sleep(1 * time.Second)
	}
}

// BinarySearchTreeDemo demonstrates BST operations
func BinarySearchTreeDemo(viz *visualizer.Visualizer) {
	tree := datastructures.NewVisualizableBinaryTree("bst_demo", viz)
	
	// Insert nodes in a specific order to create an interesting tree
	values := []interface{}{50, 30, 70, 20, 40, 60, 80, 10, 35}
	
	for _, val := range values {
		tree.Insert(val)
		time.Sleep(1 * time.Second)
	}
	
	// Perform some searches
	searchValues := []interface{}{35, 90, 20}
	for _, val := range searchValues {
		found := tree.Search(val)
		if found {
			// Value found
		}
		time.Sleep(800 * time.Millisecond)
	}
	
	// Perform in-order traversal
	tree.InOrderTraversal()
	time.Sleep(1 * time.Second)
}

// LinearSearchDemo demonstrates linear search on a list
func LinearSearchDemo(viz *visualizer.Visualizer) {
	list := datastructures.NewVisualizableList("linear_search_demo", viz)
	
	// Create a list with some data
	data := []interface{}{10, 23, 45, 12, 67, 89, 34, 56}
	for _, item := range data {
		list.Append(item)
		time.Sleep(300 * time.Millisecond)
	}
	
	// Search for a specific value
	target := 67
	found := false
	
	for i := 0; i < list.Length(); i++ {
		value, _ := list.Get(i) // This will create a "get" operation for visualization
		if value == target {
			found = true
			break
		}
		time.Sleep(600 * time.Millisecond)
	}
	
	// Add result to demonstrate completion
	if found {
		list.Append("FOUND!")
	} else {
		list.Append("NOT FOUND")
	}
}

// StackBasedExpressionEvaluationDemo demonstrates expression evaluation using a stack
func StackBasedExpressionEvaluationDemo(viz *visualizer.Visualizer) {
	stack := datastructures.NewVisualizableStack("expression_eval", viz)
	
	// Evaluate a simple postfix expression: "2 3 + 4 *" = (2 + 3) * 4 = 20
	expression := []interface{}{2, 3, "+", 4, "*"}
	
	for _, token := range expression {
		switch token {
		case "+":
			// Pop two operands
			b, _ := stack.Pop()
			time.Sleep(500 * time.Millisecond)
			a, _ := stack.Pop()
			time.Sleep(500 * time.Millisecond)
			
			// Perform operation and push result
			result := a.(int) + b.(int)
			stack.Push(result)
			time.Sleep(800 * time.Millisecond)
			
		case "*":
			// Pop two operands
			b, _ := stack.Pop()
			time.Sleep(500 * time.Millisecond)
			a, _ := stack.Pop()
			time.Sleep(500 * time.Millisecond)
			
			// Perform operation and push result
			result := a.(int) * b.(int)
			stack.Push(result)
			time.Sleep(800 * time.Millisecond)
			
		default:
			// Push operand
			stack.Push(token)
			time.Sleep(600 * time.Millisecond)
		}
	}
	
	// The final result should be on top of the stack
	result, _ := stack.Peek()
	time.Sleep(1 * time.Second)
	
	// For demonstration, we could add the result to show completion
	_ = result // Result is 20
}

// RunAllDemos runs all demonstration algorithms in sequence
func RunAllDemos(viz *visualizer.Visualizer) {
	// Run demos with delays between them
	
	go func() {
		BubbleSortDemo(viz)
	}()
	
	time.Sleep(2 * time.Second)
	
	go func() {
		StackOperationsDemo(viz)
	}()
	
	time.Sleep(3 * time.Second)
	
	go func() {
		QueueOperationsDemo(viz)
	}()
	
	time.Sleep(4 * time.Second)
	
	go func() {
		BinarySearchTreeDemo(viz)
	}()
	
	time.Sleep(5 * time.Second)
	
	go func() {
		LinearSearchDemo(viz)
	}()
	
	time.Sleep(6 * time.Second)
	
	go func() {
		StackBasedExpressionEvaluationDemo(viz)
	}()
}