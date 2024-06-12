# Monitor
[code](./monitor.go)

```
monitor ProducerConsumer {
    int count = 0;
    condition full, empty;
    void insert(Item item) {
        if (count == N) wait(full);
        putItem(item);
        count++;
        if (count == 1) signal(empty);
    }
    Item remove() {
        if (count == 0) wait(empty);
        Item item = getItem();
        count--;
        if (count == N - 1) signal(full);
        return item;
    }
}
end monitor
```