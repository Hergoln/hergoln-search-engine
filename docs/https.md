# how to run?

```mermaid
graph TD;
    client -. call .-> server;
```
server -. update and
categorize .- DB;
server -. send .- lq[log queue];