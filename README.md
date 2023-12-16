# Architecture Diagram

a workspace to work out the architecture with diagrams.

## Principles

"Diagram is also a communication tool, 'say' different things to different people."

## Tools

There are a lot of tools to draw diagrams, but I prefer to use the tools that can be used as code.

How to choose the tools?

- Easy to use (web-based, ide-integrated)
- "as Code" first, for evolvement and maintainability
- "whiteboard" is also good for quick sketching

How to draw

- Keep your style consistent
- No perfect diagram, no fancy, clear and simple is the best

### Recommended

(less is more, don't use too many tools at the same time)

- **Whiteboard** _(Not as Code)_
  - > drag-and-drop style, easy to use, but not easy to maintain.
  - Excalidraw  
  - Draw.io

- **Cloud infrastructure**
  - > Explain and visualize cloud infrastructure
  - Diagrams

- **Topology**
  - > Describe the relationship between components, e.g. network, system or interpersonal relationship
  - d2 *(with simpler DSL)*
  - Diagrams
  - ~~plantuml~~

- **Mindmap**
  - > Visualize your thoughts
  - markmap
  - mermaid
  - ~~plantuml~~ (bad style)

- **Chart**
  - > Visualize data with simple chart
  - mermaid

- **Flow/Sequence/State**
  - > Describe the process or sequence of events
  - d2 *(with simpler DSL)*
  - mermaid
  - ~~plantuml~~  

- **Object Oriented, SQL**
  - > Class, Object, Interface, or Entity-Releationship
  - d2 *(with simpler DSL)*
  - mermaid
  - ~~plantuml~~

- **Architecture Cake**
  - > Describe the layers of architecture
  - d2 *([with grid diagrams](https://d2lang.com/tour/grid-diagrams))*
  - Diagrams

- **C4 Model**
  - > C4 Model is a simple hierarchical way to think about the static structures of a software system in terms of containers, components and code.
  - mermaid
  - d2 *(not standard, but easy to use)*
  - ~~Structurizr~~
    - ~~[krzysztofreczek/go-structurizr](https://github.com/krzysztofreczek/go-structurizr)~~
    - ~~[goadesign/model](https://github.com/goadesign/model)~~

- **Terminal**
  - > terminal friendly, in case you need to draw in terminal
  - [asciiflow](https://github.com/lewish/asciiflow)

### Visualizations

- Graphviz: the base of many tools
