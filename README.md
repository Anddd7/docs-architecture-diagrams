# Architecture Diagram

a workspace to work out the architecture with diagrams.

:warning: This project is opened with CC-BY-4.0 license, you can duplicate, modify and share the materials, but you must give appropriate credit, keep the author's name and source link, and indicate if changes were made. See [CC-BY-4.0](#cc-by-40) for more details.

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

## :warning: CC-BY-4.0

- [Creative Commons Attribution 4.0 International Public License](https://creativecommons.org/licenses/by/4.0/)
- [知识共享署名 4.0 国际许可协议](https://creativecommons.org/licenses/by/4.0/deed.zh-hans)

> 根据本公共许可协议的条款，许可人授予您在全球范围内，免费的、不可再许可、非独占、不可撤销的许可，以对授权作品(Licensed Material)行使以下“协议所授予的权利”：
>
> - 复制和分享授权作品(Licensed Material)的全部或部分；以及
> - 创作、复制和分享演绎作品(Adapted Material)。

---

> 若您分享本授权作品(Licensed Material)（包含修改格式），您必须：
>
> - 保留如下标识（如果许可人提供授权作品(Licensed Material)的同时提供如下标识）：
>   - 以许可人要求的任何合理方式，标识本授权作品(Licensed Material)创作者和其他被指定署名的人的身份（包括指定的笔名）；
>   - 著作权声明；
>   - 有关本公共许可协议的声明；
>   - 有关免责的声明；
>   - 在合理可行情况下，本授权作品(Licensed Material)的网址(URI)或超链接；
> - 表明您是否修改本授权作品(Licensed Material)及保留任何先前修改的标记；及
> - 表明授权作品(Licensed Material)依据本公共许可协议授权，并提供本公共许可协议全文，或者本公共许可协议的网址(URI)或超链接。
