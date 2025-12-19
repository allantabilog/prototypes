// Visualizers for different data structures
class DataStructureVisualizer {
  constructor(container) {
    this.container = container;
    this.currentFilter = "all";
  }

  setFilter(filter) {
    this.currentFilter = filter;
    this.render();
  }

  render() {
    // Override in subclasses
  }

  clear() {
    this.container.innerHTML = "";
  }
}

class ListVisualizer extends DataStructureVisualizer {
  constructor(container) {
    super(container);
    this.lists = new Map();
  }

  updateData(snapshot) {
    this.lists.set(snapshot.id, snapshot);
    this.render();
  }

  removeData(id) {
    this.lists.delete(id);
    this.render();
  }

  render() {
    if (this.currentFilter !== "all" && this.currentFilter !== "list") {
      return;
    }

    let html = "";
    for (const [id, snapshot] of this.lists) {
      const items = snapshot.data.items || [];

      html += `
                <div class="data-structure">
                    <div class="data-structure-header">
                        List: ${id} (Length: ${snapshot.data.length})
                    </div>
                    <div class="data-structure-content">
                        <div class="list-container">
                            ${items
                              .map(
                                (item, index) =>
                                  `<div class="list-item" data-index="${index}">${item}</div>`
                              )
                              .join("")}
                            ${
                              items.length === 0
                                ? '<div class="empty-message">Empty list</div>'
                                : ""
                            }
                        </div>
                    </div>
                </div>
            `;
    }

    if (this.currentFilter === "list") {
      this.container.innerHTML =
        html || '<div class="empty-message">No lists to display</div>';
    } else {
      // Append to existing content for 'all' filter
      const listContainer = document.createElement("div");
      listContainer.innerHTML = html;
      this.container.appendChild(listContainer);
    }
  }
}

class StackVisualizer extends DataStructureVisualizer {
  constructor(container) {
    super(container);
    this.stacks = new Map();
  }

  updateData(snapshot) {
    this.stacks.set(snapshot.id, snapshot);
    this.render();
  }

  removeData(id) {
    this.stacks.delete(id);
    this.render();
  }

  render() {
    if (this.currentFilter !== "all" && this.currentFilter !== "stack") {
      return;
    }

    let html = "";
    for (const [id, snapshot] of this.stacks) {
      const items = snapshot.data.items || [];

      html += `
                <div class="data-structure">
                    <div class="data-structure-header">
                        Stack: ${id} (Size: ${snapshot.data.size})
                        ${
                          snapshot.data.top !== undefined
                            ? ` | Top: ${snapshot.data.top}`
                            : ""
                        }
                    </div>
                    <div class="data-structure-content">
                        <div class="stack-container">
                            ${items
                              .map(
                                (item, index) =>
                                  `<div class="stack-item" data-level="${index}">${item}</div>`
                              )
                              .join("")}
                            ${
                              items.length === 0
                                ? '<div class="empty-message">Empty stack</div>'
                                : ""
                            }
                        </div>
                    </div>
                </div>
            `;
    }

    if (this.currentFilter === "stack") {
      this.container.innerHTML =
        html || '<div class="empty-message">No stacks to display</div>';
    } else {
      const stackContainer = document.createElement("div");
      stackContainer.innerHTML = html;
      this.container.appendChild(stackContainer);
    }
  }
}

class QueueVisualizer extends DataStructureVisualizer {
  constructor(container) {
    super(container);
    this.queues = new Map();
  }

  updateData(snapshot) {
    this.queues.set(snapshot.id, snapshot);
    this.render();
  }

  removeData(id) {
    this.queues.delete(id);
    this.render();
  }

  render() {
    if (this.currentFilter !== "all" && this.currentFilter !== "queue") {
      return;
    }

    let html = "";
    for (const [id, snapshot] of this.queues) {
      const items = snapshot.data.items || [];

      html += `
                <div class="data-structure">
                    <div class="data-structure-header">
                        Queue: ${id} (Size: ${snapshot.data.size})
                        ${
                          snapshot.data.front !== undefined
                            ? ` | Front: ${snapshot.data.front}`
                            : ""
                        }
                        ${
                          snapshot.data.rear !== undefined
                            ? ` | Rear: ${snapshot.data.rear}`
                            : ""
                        }
                    </div>
                    <div class="data-structure-content">
                        <div class="queue-container">
                            ${items
                              .map(
                                (item, index) =>
                                  `<div class="queue-item" data-position="${index}">${item}</div>`
                              )
                              .join("")}
                            ${
                              items.length === 0
                                ? '<div class="empty-message">Empty queue</div>'
                                : ""
                            }
                        </div>
                    </div>
                </div>
            `;
    }

    if (this.currentFilter === "queue") {
      this.container.innerHTML =
        html || '<div class="empty-message">No queues to display</div>';
    } else {
      const queueContainer = document.createElement("div");
      queueContainer.innerHTML = html;
      this.container.appendChild(queueContainer);
    }
  }
}

class TreeVisualizer extends DataStructureVisualizer {
  constructor(container) {
    super(container);
    this.trees = new Map();
    this.nodeRadius = 20;
    this.levelHeight = 60;
  }

  updateData(snapshot) {
    this.trees.set(snapshot.id, snapshot);
    this.render();
  }

  removeData(id) {
    this.trees.delete(id);
    this.render();
  }

  render() {
    if (this.currentFilter !== "all" && this.currentFilter !== "binary_tree") {
      return;
    }

    let html = "";
    for (const [id, snapshot] of this.trees) {
      html += `
                <div class="data-structure">
                    <div class="data-structure-header">
                        Binary Tree: ${id} (Size: ${
        snapshot.data.size
      }, Height: ${snapshot.metadata.height})
                    </div>
                    <div class="data-structure-content">
                        <div class="tree-container" id="tree-${id}">
                            ${
                              snapshot.data.size === 0
                                ? '<div class="empty-message">Empty tree</div>'
                                : ""
                            }
                        </div>
                    </div>
                </div>
            `;
    }

    if (this.currentFilter === "binary_tree") {
      this.container.innerHTML =
        html || '<div class="empty-message">No trees to display</div>';
    } else {
      const treeContainer = document.createElement("div");
      treeContainer.innerHTML = html;
      this.container.appendChild(treeContainer);
    }

    // Render tree visualizations using D3
    for (const [id, snapshot] of this.trees) {
      if (snapshot.data.root) {
        this.renderTree(id, snapshot.data.root);
      }
    }
  }

  renderTree(treeId, rootData) {
    const container = document.getElementById(`tree-${treeId}`);
    if (!container) return;

    // Clear any existing SVG
    d3.select(container).select("svg").remove();

    // Create hierarchy from data
    const root = d3.hierarchy(rootData, (d) => {
      const children = [];
      if (d.left) children.push(d.left);
      if (d.right) children.push(d.right);
      return children.length ? children : null;
    });

    // Calculate tree layout
    const treeLayout = d3
      .tree()
      .size([400, 200])
      .separation((a, b) => 1);

    treeLayout(root);

    // Create SVG
    const svg = d3
      .select(container)
      .append("svg")
      .attr("width", 500)
      .attr("height", 250)
      .style("margin", "1rem auto")
      .style("display", "block");

    const g = svg.append("g").attr("transform", "translate(50, 25)");

    // Create links
    const links = g
      .selectAll(".tree-link")
      .data(root.links())
      .enter()
      .append("line")
      .attr("class", "tree-link")
      .attr("x1", (d) => d.source.x)
      .attr("y1", (d) => d.source.y)
      .attr("x2", (d) => d.target.x)
      .attr("y2", (d) => d.target.y);

    // Create nodes
    const nodes = g
      .selectAll(".tree-node-group")
      .data(root.descendants())
      .enter()
      .append("g")
      .attr("class", "tree-node-group")
      .attr("transform", (d) => `translate(${d.x}, ${d.y})`);

    nodes
      .append("circle")
      .attr("class", "tree-node")
      .attr("r", this.nodeRadius);

    nodes
      .append("text")
      .attr("class", "tree-node-text")
      .text((d) => d.data.value);
  }
}

// Combined visualizer manager
class VisualizerManager {
  constructor(container) {
    this.container = container;
    this.listViz = new ListVisualizer(container);
    this.stackViz = new StackVisualizer(container);
    this.queueViz = new QueueVisualizer(container);
    this.treeViz = new TreeVisualizer(container);
    this.currentFilter = "all";
  }

  setFilter(filter) {
    this.currentFilter = filter;
    this.container.innerHTML = "";

    this.listViz.setFilter(filter);
    this.stackViz.setFilter(filter);
    this.queueViz.setFilter(filter);
    this.treeViz.setFilter(filter);
  }

  updateSnapshot(snapshot) {
    switch (snapshot.type) {
      case "list":
        this.listViz.updateData(snapshot);
        break;
      case "stack":
        this.stackViz.updateData(snapshot);
        break;
      case "queue":
        this.queueViz.updateData(snapshot);
        break;
      case "binary_tree":
        this.treeViz.updateData(snapshot);
        break;
    }
  }

  removeDataStructure(id, type) {
    switch (type) {
      case "list":
        this.listViz.removeData(id);
        break;
      case "stack":
        this.stackViz.removeData(id);
        break;
      case "queue":
        this.queueViz.removeData(id);
        break;
      case "binary_tree":
        this.treeViz.removeData(id);
        break;
    }
  }

  clear() {
    this.container.innerHTML = "";
    this.listViz.lists.clear();
    this.stackViz.stacks.clear();
    this.queueViz.queues.clear();
    this.treeViz.trees.clear();
  }
}
