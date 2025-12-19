package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/allantabilog/visualiser/examples"
	"github.com/allantabilog/visualiser/internal/datastructures"
	"github.com/allantabilog/visualiser/internal/server"
	"github.com/allantabilog/visualiser/internal/visualizer"
)

// InteractiveDemo provides an interactive command-line interface for the visualizer
type InteractiveDemo struct {
	viz           *visualizer.Visualizer
	dataStructures map[string]interface{}
	scanner       *bufio.Scanner
}

// NewInteractiveDemo creates a new interactive demo
func NewInteractiveDemo(viz *visualizer.Visualizer) *InteractiveDemo {
	return &InteractiveDemo{
		viz:           viz,
		dataStructures: make(map[string]interface{}),
		scanner:       bufio.NewScanner(os.Stdin),
	}
}

// Start begins the interactive demo
func (demo *InteractiveDemo) Start() {
	fmt.Println("=== Data Structure Visualizer Interactive Demo ===")
	fmt.Println("Open http://localhost:8081 in your browser to see the visualizations")
	fmt.Println()
	
	for {
		demo.showMenu()
		choice := demo.getInput("Enter your choice: ")
		
		switch choice {
		case "1":
			demo.createList()
		case "2":
			demo.createStack()
		case "3":
			demo.createQueue()
		case "4":
			demo.createBinaryTree()
		case "5":
			demo.runBubbleSort()
		case "6":
			demo.runStackOperations()
		case "7":
			demo.runQueueOperations()
		case "8":
			demo.runBinaryTreeOperations()
		case "9":
			demo.runLinearSearch()
		case "10":
			demo.runAllDemos()
		case "11":
			demo.listDataStructures()
		case "12":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		
		fmt.Println()
	}
}

func (demo *InteractiveDemo) showMenu() {
	fmt.Println("=== Main Menu ===")
	fmt.Println("Data Structure Creation:")
	fmt.Println("  1. Create List")
	fmt.Println("  2. Create Stack")
	fmt.Println("  3. Create Queue")
	fmt.Println("  4. Create Binary Tree")
	fmt.Println()
	fmt.Println("Algorithm Demonstrations:")
	fmt.Println("  5. Bubble Sort Demo")
	fmt.Println("  6. Stack Operations Demo")
	fmt.Println("  7. Queue Operations Demo")
	fmt.Println("  8. Binary Tree Operations Demo")
	fmt.Println("  9. Linear Search Demo")
	fmt.Println("  10. Run All Demos")
	fmt.Println()
	fmt.Println("Other:")
	fmt.Println("  11. List Active Data Structures")
	fmt.Println("  12. Exit")
	fmt.Println()
}

func (demo *InteractiveDemo) getInput(prompt string) string {
	fmt.Print(prompt)
	demo.scanner.Scan()
	return strings.TrimSpace(demo.scanner.Text())
}

func (demo *InteractiveDemo) createList() {
	id := demo.getInput("Enter list ID: ")
	if _, exists := demo.dataStructures[id]; exists {
		fmt.Printf("Data structure with ID '%s' already exists.\n", id)
		return
	}
	
	list := datastructures.NewVisualizableList(id, demo.viz)
	demo.dataStructures[id] = list
	
	fmt.Printf("Created list '%s'. Adding some sample data...\n", id)
	
	// Add some sample data
	for i := 1; i <= 5; i++ {
		list.Append(i * 10)
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Printf("List '%s' created with sample data: [10, 20, 30, 40, 50]\n", id)
}

func (demo *InteractiveDemo) createStack() {
	id := demo.getInput("Enter stack ID: ")
	if _, exists := demo.dataStructures[id]; exists {
		fmt.Printf("Data structure with ID '%s' already exists.\n", id)
		return
	}
	
	stack := datastructures.NewVisualizableStack(id, demo.viz)
	demo.dataStructures[id] = stack
	
	fmt.Printf("Created stack '%s'. Adding some sample data...\n", id)
	
	// Add some sample data
	items := []string{"A", "B", "C", "D", "E"}
	for _, item := range items {
		stack.Push(item)
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Printf("Stack '%s' created with sample data\n", id)
}

func (demo *InteractiveDemo) createQueue() {
	id := demo.getInput("Enter queue ID: ")
	if _, exists := demo.dataStructures[id]; exists {
		fmt.Printf("Data structure with ID '%s' already exists.\n", id)
		return
	}
	
	queue := datastructures.NewVisualizableQueue(id, demo.viz)
	demo.dataStructures[id] = queue
	
	fmt.Printf("Created queue '%s'. Adding some sample data...\n", id)
	
	// Add some sample data
	for i := 1; i <= 5; i++ {
		queue.Enqueue(fmt.Sprintf("Item-%d", i))
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Printf("Queue '%s' created with sample data\n", id)
}

func (demo *InteractiveDemo) createBinaryTree() {
	id := demo.getInput("Enter binary tree ID: ")
	if _, exists := demo.dataStructures[id]; exists {
		fmt.Printf("Data structure with ID '%s' already exists.\n", id)
		return
	}
	
	tree := datastructures.NewVisualizableBinaryTree(id, demo.viz)
	demo.dataStructures[id] = tree
	
	fmt.Printf("Created binary tree '%s'. Adding some sample data...\n", id)
	
	// Add some sample data
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range values {
		tree.Insert(val)
		time.Sleep(800 * time.Millisecond)
	}
	
	fmt.Printf("Binary tree '%s' created with sample data\n", id)
}

func (demo *InteractiveDemo) runBubbleSort() {
	fmt.Println("Running Bubble Sort demonstration...")
	go examples.BubbleSortDemo(demo.viz)
	fmt.Println("Bubble sort demo started. Check the web interface for visualization.")
}

func (demo *InteractiveDemo) runStackOperations() {
	fmt.Println("Running Stack Operations demonstration...")
	go examples.StackOperationsDemo(demo.viz)
	fmt.Println("Stack operations demo started. Check the web interface for visualization.")
}

func (demo *InteractiveDemo) runQueueOperations() {
	fmt.Println("Running Queue Operations demonstration...")
	go examples.QueueOperationsDemo(demo.viz)
	fmt.Println("Queue operations demo started. Check the web interface for visualization.")
}

func (demo *InteractiveDemo) runBinaryTreeOperations() {
	fmt.Println("Running Binary Tree Operations demonstration...")
	go examples.BinarySearchTreeDemo(demo.viz)
	fmt.Println("Binary tree demo started. Check the web interface for visualization.")
}

func (demo *InteractiveDemo) runLinearSearch() {
	fmt.Println("Running Linear Search demonstration...")
	go examples.LinearSearchDemo(demo.viz)
	fmt.Println("Linear search demo started. Check the web interface for visualization.")
}

func (demo *InteractiveDemo) runAllDemos() {
	fmt.Println("Running all demonstrations...")
	go examples.RunAllDemos(demo.viz)
	fmt.Println("All demos started. Check the web interface for visualizations.")
}

func (demo *InteractiveDemo) listDataStructures() {
	if len(demo.dataStructures) == 0 {
		fmt.Println("No data structures created yet.")
		return
	}
	
	fmt.Println("Active data structures:")
	for id, ds := range demo.dataStructures {
		switch ds.(type) {
		case *datastructures.VisualizableList:
			fmt.Printf("  - %s (List)\n", id)
		case *datastructures.VisualizableStack:
			fmt.Printf("  - %s (Stack)\n", id)
		case *datastructures.VisualizableQueue:
			fmt.Printf("  - %s (Queue)\n", id)
		case *datastructures.VisualizableBinaryTree:
			fmt.Printf("  - %s (Binary Tree)\n", id)
		}
	}
}

func main() {
	// Create the visualizer
	viz := visualizer.NewVisualizer()
	
	// Create and start the server in a goroutine
	srv := server.NewServer(viz)
	
	go func() {
		log.Printf("Data Structure Visualizer starting...")
		log.Printf("Server will be available at http://localhost:8081")
		if err := srv.Start(8081); err != nil {
			log.Printf("Server failed to start: %v", err)
			log.Printf("You can still use the demo, but web interface won't be available")
		}
	}()
	
	// Give the server a moment to start
	time.Sleep(2 * time.Second)
	
	// Start interactive demo
	demo := NewInteractiveDemo(viz)
	demo.Start()
}