# Mastering Go

## Code organization on Gin Framework
- No mandatory organization
- MVC Architecture

    ```mermaid
    flowchart LR


    R[User] -->|Request| C[controller]
    C -->|Request Data| M[Model]
    M -->|Data| C
    C -->|Data| V[View]
    V -->|Response| R
    M --> D[(Database)]
    ```
```shel

│   ├── controllers
│   ├── helpers
│   ├── middlewares
│   ├── models
│   ├── static
│   └── templates
│   ├── go.mod
│   ├── main.go
|___README.md
```