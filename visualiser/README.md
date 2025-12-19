# Data Structure Visualizer

A comprehensive real-time visualization tool for data structures and algorithms implemented in Go, with an interactive web-based frontend using JavaScript and D3.js.

## 🌟 Features

- **Real-time visualization** of data structures (lists, stacks, queues, binary trees)
- **Snapshot capability** to capture state at any point in time
- **Web-based interface** using modern JavaScript libraries and D3.js for tree visualization
- **WebSocket communication** for live updates and real-time synchronization
- **Interactive controls** with tabbed interface for different data structure types
- **Operation logging** with timestamps and parameter tracking
- **Example algorithms** including bubble sort, linear search, and tree operations
- **Command-line demos** for easy testing and demonstration

## 🏗️ Architecture

```
├── cmd/                    # Application entry points
│   ├── visualiser/         # Main visualiser server application
│   └── demo/               # Interactive demo application
├── internal/               # Private application code
│   ├── datastructures/     # Visualizable data structure implementations
│   │   ├── list.go         # Dynamic array with visualization
│   │   ├── stack.go        # LIFO stack with visualization
│   │   ├── queue.go        # FIFO queue with visualization
│   │   └── binary_tree.go  # Binary search tree with visualization
│   ├── visualizer/         # Core visualization framework
│   │   ├── visualizer.go   # Main visualization coordinator
│   │   └── errors.go       # Error definitions
│   └── server/             # Web server and WebSocket handlers
│       └── server.go       # HTTP server with WebSocket support
├── web/                    # Web frontend files
│   ├── static/            # CSS, JS, and client-side assets
│   │   ├── styles.css     # Comprehensive styling for all visualizations
│   │   ├── websocket.js   # WebSocket client management
│   │   ├── visualizers.js # D3.js-based data structure renderers
│   │   └── app.js         # Main application controller
│   └── templates/         # HTML templates
│       └── index.html     # Main application interface
├── examples/              # Example algorithms and demonstrations
│   └── algorithms.go      # Collection of algorithm demonstrations
└── docs/                  # Documentation (future expansion)
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- A modern web browser with JavaScript enabled

### Installation

1. Clone or create the project:

```bash
mkdir -p ~/dev/golang/prototypes/visualiser
cd ~/dev/golang/prototypes/visualiser
```

2. Initialize the Go module and install dependencies:

```bash
go mod init github.com/allantabilog/visualiser
go mod tidy
```

### Running the Application

#### Option 1: Basic Server

Start the visualization server:

```bash
go run cmd/visualiser/main.go
```

Then open your browser to `http://localhost:8080` to see the web interface.

#### Option 2: Server with Auto-Demo

Run the server with automatic algorithm demonstrations:

```bash
go run cmd/visualiser/main.go -demo
```

#### Option 3: Interactive Demo

Run the interactive command-line demo:

```bash
go run cmd/demo/main.go
```

This provides a menu-driven interface where you can:

- Create individual data structures
- Run specific algorithm demonstrations
- Interact with the visualizations through the command line

### Command Line Options

For the main visualizer server (`cmd/visualiser/main.go`):

- `-demo`: Run demonstration algorithms automatically
- `-port <number>`: Specify the port (default: 8080)

## 📊 Supported Data Structures

### 1. Lists (Dynamic Arrays)

- **Operations**: Append, Insert, Delete, Get, Clear
- **Visualization**: Horizontal array with indexed elements
- **Use Cases**: General-purpose storage, algorithm demonstrations

### 2. Stacks (LIFO)

- **Operations**: Push, Pop, Peek, Clear
- **Visualization**: Vertical stack with top element highlighted
- **Use Cases**: Expression evaluation, function calls, undo operations

### 3. Queues (FIFO)

- **Operations**: Enqueue, Dequeue, Front, Rear, Clear
- **Visualization**: Horizontal queue with front/rear indicators
- **Use Cases**: Task scheduling, breadth-first search, buffer management

### 4. Binary Trees

- **Operations**: Insert, Search, Traversal, Clear
- **Visualization**: Hierarchical tree structure using D3.js
- **Use Cases**: Searching, sorting, hierarchical data representation

## 🎯 Example Algorithms

The system includes several pre-built algorithm demonstrations:

### Sorting Algorithms

- **Bubble Sort**: Visual step-by-step sorting with element swapping
- Demonstrates comparison-based sorting with real-time updates

### Search Algorithms

- **Linear Search**: Sequential search through list elements
- **Binary Search Tree Operations**: Insert, search, and traversal

### Data Structure Operations

- **Stack-based Expression Evaluation**: Postfix expression calculator
- **Queue Simulation**: Customer service queue management
- **Tree Traversal**: In-order, pre-order, and post-order traversals

## 🖥️ Web Interface Features

### Real-time Visualization

- **Live Updates**: All operations are immediately reflected in the browser
- **Multiple Views**: Tabbed interface for different data structure types
- **Operation Logging**: Real-time log of all operations with timestamps

### Interactive Elements

- **Data Structure List**: Shows all active data structures with size information
- **Operation History**: Chronological list of recent operations
- **Filter Tabs**: View all structures or filter by type
- **Connection Status**: Real-time WebSocket connection monitoring

### Responsive Design

- **Mobile-friendly**: Adapts to different screen sizes
- **Modern UI**: Clean, professional interface with smooth animations
- **Accessibility**: Keyboard navigation and screen reader support

## 🔧 API Reference

### WebSocket Messages

#### Client to Server

```javascript
// Request a specific snapshot
{
  "type": "get_snapshot",
  "id": "data_structure_id"
}
```

#### Server to Client

```javascript
// Initial state
{
  "type": "initial_state",
  "data": {
    "id1": { /* snapshot */ },
    "id2": { /* snapshot */ }
  }
}

// Operation update
{
  "type": "operation",
  "data": {
    "id": "operation_id",
    "dataStructure": "structure_id",
    "type": "insert",
    "parameters": { "item": "value" },
    "timestamp": "2023-...",
    "before": { /* snapshot */ },
    "after": { /* snapshot */ }
  }
}
```

### REST API Endpoints

- `GET /api/snapshots` - Get all current snapshots
- `GET /api/snapshots/{id}` - Get specific snapshot
- `GET /api/operations` - Get all operations
- `GET /api/operations/{id}` - Get operations for specific data structure

## 🧪 Testing and Development

### Running Individual Components

1. **Test the visualizer core**:

```go
viz := visualizer.NewVisualizer()
list := datastructures.NewVisualizableList("test", viz)
list.Append(42)
snapshot, _ := list.Snapshot()
fmt.Printf("Snapshot: %+v\n", snapshot)
```

2. **Test WebSocket connectivity**:
   Open the browser developer console and check for WebSocket connection messages.

3. **Test REST API**:

```bash
curl http://localhost:8080/api/snapshots
```

### Extending the System

#### Adding New Data Structures

1. Implement the `Visualizable` interface
2. Add visualization logic in `web/static/visualizers.js`
3. Update CSS styles in `web/static/styles.css`

#### Adding New Algorithms

1. Create algorithm functions in `examples/algorithms.go`
2. Add menu options in the demo application
3. Test with the interactive demo

## 🎨 Customization

### Styling

Modify `web/static/styles.css` to customize:

- Color schemes for different data structures
- Animation timings and effects
- Layout and spacing
- Responsive breakpoints

### Visualization Behavior

Adjust `web/static/visualizers.js` to change:

- Animation speeds
- Display formats
- Interactive behaviors
- D3.js tree layout parameters

## 🐛 Troubleshooting

### Common Issues

1. **WebSocket Connection Failed**

   - Check if port 8080 is available
   - Verify firewall settings
   - Try a different port with `-port` flag

2. **Browser Not Updating**

   - Check browser console for JavaScript errors
   - Ensure WebSocket connection is established
   - Try refreshing the page

3. **Go Module Issues**
   - Run `go mod tidy` to resolve dependencies
   - Check Go version (requires 1.21+)

### Performance Considerations

- The system is designed for educational/demonstration purposes
- For large datasets, consider adjusting animation delays
- WebSocket message size is limited; very large trees may need optimization

## 🔮 Future Enhancements

### Planned Features

- **More Data Structures**: Hash tables, heaps, graphs, tries
- **Advanced Algorithms**: Merge sort, quick sort, graph algorithms
- **Export Capabilities**: Save visualizations as images or videos
- **Step-by-step Debugging**: Pause and step through algorithms
- **Code Integration**: Show actual code alongside visualizations

### Architecture Improvements

- **Plugin System**: Easy addition of new data structures
- **Persistence**: Save and load visualization sessions
- **Multi-user Support**: Collaborative visualization sessions
- **Performance Monitoring**: Track algorithm complexity in real-time

## 📝 Contributing

This is a prototype project demonstrating data structure visualization concepts. To extend or modify:

1. Fork the repository
2. Create feature branches for new data structures or algorithms
3. Test thoroughly with the demo application
4. Ensure responsive web interface works across browsers
5. Update documentation for new features

## 📄 License

This project is intended for educational and demonstration purposes. Feel free to use and modify as needed for learning about data structures, algorithms, and real-time web applications.

## 🙏 Acknowledgments

- **D3.js** for powerful tree visualization capabilities
- **Gorilla WebSocket** for reliable real-time communication
- **Go community** for excellent standard library and ecosystem
