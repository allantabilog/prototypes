package main

import (
	"time"

	"github.com/allantabilog/visualiser/internal/datastructures"
	"github.com/allantabilog/visualiser/internal/server"
	"github.com/allantabilog/visualiser/internal/visualizer"
)

// Simple example showing how to create and use visualizable data structures
func main() {
	// Create the visualizer
	viz := visualizer.NewVisualizer()
	
	// Start the web server in background
	srv := server.NewServer(viz)
	go srv.Start(8080)
	
	// Give server time to start
	time.Sleep(1 * time.Second)
	println("Visualizer running at http://localhost:8080")
	
	// Create a list and add some data
	list := datastructures.NewVisualizableList("my_list", viz)
	for i := 1; i <= 5; i++ {
		list.Append(i * 10)
		time.Sleep(1 * time.Second)
	}
	
	// Create a stack and demonstrate operations
	stack := datastructures.NewVisualizableStack("my_stack", viz)
	stack.Push("First")
	time.Sleep(1 * time.Second)
	stack.Push("Second")
	time.Sleep(1 * time.Second)
	stack.Push("Third")
	time.Sleep(1 * time.Second)
	
	popped, _ := stack.Pop()
	println("Popped:", popped.(string))
	time.Sleep(1 * time.Second)
	
	// Create a binary tree
	tree := datastructures.NewVisualizableBinaryTree("my_tree", viz)
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		tree.Insert(val)
		time.Sleep(1 * time.Second)
	}
	
	// Keep the program running
	println("Data structures created! Check the web interface.")
	println("Press Ctrl+C to exit")
	select {} // Block forever
}