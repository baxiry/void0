# 

 is a simple command line to interactive with zaradb, it make end to end testing zaradb so easy.



## Usage


```
$ zarashell ws://localhost:3000/ws
> {"type": "echo", "payload": "Hello, world"}
< {"type":"echo","payload":"Hello, world"}
> {"type": "broadcast", "payload": "Hello, world"}
< {"type":"broadcast","payload":"Hello, world"}
< {"type":"broadcastResult","payload":"Hello, world","listenerCount":1}
> ^D
EOF
```
