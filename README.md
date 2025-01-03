To generate plantuml diagram from OpenAPI & Arazzo files, you can use the following command:

### Run with Go

If you have go >= 1.16 installed, you can run the following command to build and run flowline:

```bash title="Install flowline"
go run github.com/theapix/flowline@latest
```

### Quick Start

```bash title="Generate PlantUML diagram for all Arazzo worflows"
flowline --arazzo <path-to-arazzo-file>
```

```bash title="Generate PlantUML diagram for a specific Arazzo workflow"
flowline --arazzo <path-to-arazzo-file> --workflow <workflow-id>
```
