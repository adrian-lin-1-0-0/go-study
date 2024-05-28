# HTTP over Unix Domain Sockets

### Server

直接把`net.Listener`注入`http.Serve` 即可

```go
l, err := net.Listen("unix", sockFileName)
if err != nil {
    log.Fatal("Listen error: ", err)
}
defer l.Close()


http.Serve(l,
    ...,
)
```

### Client

注入`http.Client.Transport`,直接無視`Dial`的`proto`和`addr`，直接返回`net.Dial("unix", sockFileName)`

```go
client := &http.Client{
    Transport: &http.Transport{
        Dial: func(proto, addr string) (conn net.Conn, err error) {
            return net.Dial("unix", sockFileName)
        },
    },
}

resp, err := client.Get("http://xxx/")
```