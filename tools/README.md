# Veeva MDL Tools

## Mermaid ERD Generator
`generate_erd.py` parses the `mdl_components_schema.json` and generates a [Mermaid](https://mermaid.js.org/) Class Diagram.

### Files
- `veeva_mdl_erd.mmd`: The generated Mermaid definition file.

### How to Generate Image
Since the diagram is large (~2000 lines), it is best rendered using the Mermaid CLI (`mmdc`).

**Prerequisites:**
- Node.js & npm

**Install Mermaid CLI:**
```bash
npm install -g @mermaid-js/mermaid-cli
```

**Render to PNG:**
```bash
mmdc -i veeva_mdl_erd.mmd -o veeva_mdl_erd.png --width 4000 --height 4000
```

**Render to SVG:**
```bash
mmdc -i veeva_mdl_erd.mmd -o veeva_mdl_erd.svg
```
