// Main application controller
class VisualizerApp {
  constructor() {
    this.wsManager = new WebSocketManager();
    this.visualizerManager = null;
    this.dataStructures = new Map();
    this.operations = [];

    this.initializeElements();
    this.setupEventListeners();
    this.initializeVisualizers();
    this.connectWebSocket();
  }

  initializeElements() {
    this.statusElement = document.getElementById("status");
    this.clearBtn = document.getElementById("clear-btn");
    this.reconnectBtn = document.getElementById("reconnect-btn");
    this.visualizationContainer = document.getElementById(
      "visualization-container"
    );
    this.dataStructuresList = document.getElementById("data-structures-list");
    this.operationsLog = document.getElementById("operations-log");
    this.tabButtons = document.querySelectorAll(".tab-btn");
  }

  setupEventListeners() {
    // Clear button
    this.clearBtn.addEventListener("click", () => {
      this.clearAll();
    });

    // Reconnect button
    this.reconnectBtn.addEventListener("click", () => {
      this.wsManager.disconnect();
      this.wsManager.connect();
    });

    // Tab buttons
    this.tabButtons.forEach((btn) => {
      btn.addEventListener("click", () => {
        this.setActiveTab(btn);
        const filter = btn.dataset.tab;
        this.visualizerManager.setFilter(filter);
      });
    });
  }

  initializeVisualizers() {
    this.visualizerManager = new VisualizerManager(this.visualizationContainer);
  }

  connectWebSocket() {
    this.wsManager.onMessage = (message) => {
      this.handleWebSocketMessage(message);
    };

    this.wsManager.onStatusChange = (status) => {
      this.updateConnectionStatus(status);
    };

    this.wsManager.connect();
  }

  handleWebSocketMessage(message) {
    console.log("Received message:", message);

    switch (message.type) {
      case "initial_state":
        this.handleInitialState(message.data);
        break;
      case "operation":
        this.handleOperation(message.data);
        break;
      case "snapshot":
        this.handleSnapshot(message.data);
        break;
      default:
        console.log("Unknown message type:", message.type);
    }
  }

  handleInitialState(snapshots) {
    console.log("Received initial state:", snapshots);

    // Clear existing data
    this.dataStructures.clear();
    this.operations = [];

    // Update with initial data
    for (const [id, snapshot] of Object.entries(snapshots)) {
      this.dataStructures.set(id, snapshot);
      this.visualizerManager.updateSnapshot(snapshot);
    }

    this.updateDataStructuresList();
    this.updateOperationsLog();
  }

  handleOperation(operation) {
    console.log("Received operation:", operation);

    // Add to operations log
    this.operations.unshift(operation);
    if (this.operations.length > 100) {
      this.operations = this.operations.slice(0, 100);
    }

    // Update visualization if we have the after snapshot
    if (operation.after) {
      this.dataStructures.set(operation.dataStructure, operation.after);
      this.visualizerManager.updateSnapshot(operation.after);
    }

    this.updateDataStructuresList();
    this.updateOperationsLog();

    // Highlight the operation briefly
    this.highlightRecentOperation(operation);
  }

  handleSnapshot(snapshot) {
    console.log("Received snapshot:", snapshot);

    this.dataStructures.set(snapshot.id, snapshot);
    this.visualizerManager.updateSnapshot(snapshot);
    this.updateDataStructuresList();
  }

  updateDataStructuresList() {
    let html = "";

    for (const [id, snapshot] of this.dataStructures) {
      const sizeInfo = this.getSizeInfo(snapshot);
      html += `
                <div class="ds-item" data-id="${id}" data-type="${
        snapshot.type
      }">
                    <div class="ds-type">${this.formatType(snapshot.type)}</div>
                    <div class="ds-id">${id}</div>
                    <div class="ds-size">${sizeInfo}</div>
                </div>
            `;
    }

    this.dataStructuresList.innerHTML =
      html || '<div class="empty-message">No data structures</div>';

    // Add click listeners to data structure items
    this.dataStructuresList.querySelectorAll(".ds-item").forEach((item) => {
      item.addEventListener("click", () => {
        const id = item.dataset.id;
        this.wsManager.requestSnapshot(id);
      });
    });
  }

  updateOperationsLog() {
    let html = "";

    for (const operation of this.operations.slice(0, 20)) {
      const timeStr = new Date(operation.timestamp).toLocaleTimeString();
      const isRecent =
        Date.now() - new Date(operation.timestamp).getTime() < 5000;

      html += `
                <div class="operation-item ${isRecent ? "recent" : ""}">
                    <span class="operation-type">${operation.type}</span>
                    <span class="operation-time">${timeStr}</span>
                    <div class="operation-target">${
                      operation.dataStructure
                    }</div>
                    ${this.formatOperationParams(operation.parameters)}
                </div>
            `;
    }

    this.operationsLog.innerHTML =
      html || '<div class="empty-message">No operations yet</div>';
  }

  getSizeInfo(snapshot) {
    switch (snapshot.type) {
      case "list":
        return `Length: ${snapshot.data.length || 0}`;
      case "stack":
      case "queue":
        return `Size: ${snapshot.data.size || 0}`;
      case "binary_tree":
        return `Size: ${snapshot.data.size || 0}, Height: ${
          snapshot.metadata?.height || 0
        }`;
      default:
        return "Unknown";
    }
  }

  formatType(type) {
    switch (type) {
      case "binary_tree":
        return "Binary Tree";
      default:
        return type.charAt(0).toUpperCase() + type.slice(1);
    }
  }

  formatOperationParams(params) {
    if (!params || Object.keys(params).length === 0) {
      return "";
    }

    const formatted = Object.entries(params)
      .map(([key, value]) => `${key}: ${value}`)
      .join(", ");

    return `<div class="operation-params">${formatted}</div>`;
  }

  setActiveTab(activeBtn) {
    this.tabButtons.forEach((btn) => btn.classList.remove("active"));
    activeBtn.classList.add("active");
  }

  highlightRecentOperation(operation) {
    // Add visual feedback for recent operations
    const dsItem = this.dataStructuresList.querySelector(
      `[data-id="${operation.dataStructure}"]`
    );
    if (dsItem) {
      dsItem.classList.add("recent-activity");
      setTimeout(() => {
        dsItem.classList.remove("recent-activity");
      }, 2000);
    }
  }

  updateConnectionStatus(status) {
    this.statusElement.textContent = status;

    // Update UI based on connection status
    const isConnected = status === "Connected";
    this.reconnectBtn.disabled = isConnected;

    if (!isConnected) {
      this.visualizationContainer.innerHTML =
        '<div class="connection-message">Disconnected from server. Click reconnect to try again.</div>';
    }
  }

  clearAll() {
    this.dataStructures.clear();
    this.operations = [];
    this.visualizerManager.clear();
    this.updateDataStructuresList();
    this.updateOperationsLog();
  }

  // Utility method to request specific snapshots
  requestSnapshot(id) {
    this.wsManager.requestSnapshot(id);
  }
}

// Initialize the application when the DOM is loaded
document.addEventListener("DOMContentLoaded", () => {
  window.visualizerApp = new VisualizerApp();
});

// Add some additional CSS for dynamic styles
const style = document.createElement("style");
style.textContent = `
    .recent-activity {
        animation: pulse 1s ease-in-out;
    }
    
    @keyframes pulse {
        0% { background-color: #f8f9fa; }
        50% { background-color: #fff3cd; }
        100% { background-color: #f8f9fa; }
    }
    
    .connection-message {
        text-align: center;
        padding: 2rem;
        color: #666;
        background: #f8f9fa;
        border-radius: 8px;
        margin: 2rem;
    }
    
    .empty-message {
        text-align: center;
        padding: 1rem;
        color: #666;
        font-style: italic;
    }
    
    .operation-params {
        font-size: 0.8rem;
        color: #666;
        margin-top: 0.25rem;
    }
    
    .ds-size {
        font-size: 0.8rem;
        color: #666;
    }
`;
document.head.appendChild(style);
