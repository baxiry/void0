# ws

zarashell is a simple command line to interactive with zaradb, it make end to end testing zaradb so easy.

![Example usage recording](https://hashrocket-production.s3.amazonaws.com/uploads/blog/misc/ws/ws.gif)


## Usage

Simply run 'zarashell' command with the your zaradb URL

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
