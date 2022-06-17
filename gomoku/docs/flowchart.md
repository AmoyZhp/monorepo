```mermaid
flowchart TD
    A([ Start ]) --> B[Move Info]
    B --> C[Chessboard set move]
    C --> D[SearchEngine take chessboard to find next move]
    D --> E[return next move]
    E --> F([END])
```