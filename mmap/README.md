# Memory Map a File (MMAP)

```sh
echo hello > /tmp/mmap-test 
```

```go
at, err := mmap.Open("/tmp/mmap-test")

if err != nil {
    log.Fatal(err)
}

buff := make([]byte, 5)
_, _ = at.ReadAt(buff, 0)
_ = at.Close()
fmt.Println(string(buff))
```

output:
```
hello
```