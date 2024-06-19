# go-study
> OS的部分大部分有golang實現

## golang
- sync/atomic
- golang.org/x/sync/semaphore

## Pattern
- [singleflight](./singleflight/README.md)
- [clean architecture]
    - [transaction](./clean_architecture/README.md)

## OS
- [unix domain socket](./unix_domain_socket/README.md)
    - [http over unix domain socket](./unix_domain_socket/http_over_unix_socket/README.md)
- [mmap](./mmap/README.md)
- [redlock](https://github.com/adrian-lin-1-0-0/redlock)
- [assembler](https://github.com/adrian-lin-1-0-0/assembly-go)

- [Process Synchronization](./os/process_synchronization/README.md)
    - [Deadlock](./os/process_synchronization/deadlock/README.md)
    - [Peterson Algorithm](./os/process_synchronization/peterson/README.md)
    - [Mutex Lock](./os/process_synchronization/lock/README.md)
        - [ ] golang mutex
    - [Compare And Swap](./os/process_synchronization/compare_and_swap/README.md)
    - [Swap](./os/process_synchronization/swap/README.md)
    - [Test and Set](./os/process_synchronization/test_and_set/README.md)
    - [Semaphore](./os/process_synchronization/semaphore/README.md)
        - [ ] [golang.org/x/sync/semaphore]
    - [Dining Philosophers Problem](./os/process_synchronization/dining_philosophers_problem/README.md)
        - [ ] [Dijkstra's solution]
        - [ ] [Chandy/Misra solution]
- [ ] [Communicating Sequential Processes (CSP)](./communicating_sequential_processes/README.md)
    - [ ] [Synchronous Communication](./communicating_sequential_processes/synchronous_communication/README.md)
    - [Producter-Consumer Problem](./os/process_synchronization/semaphore/producer-consumer_problem/README.md)
    - [Monitor](./os/process_synchronization/monitor/README.md)
- [Linux 源碼分析](https://github.com/liexusong/linux-source-code-analyze/tree/master?tab=readme-ov-file)
    - [x] [unix domain socket](https://github.com/liexusong/linux-source-code-analyze/blob/master/unix-domain-sockets.md)
    - [ ] zero copy


## Tools

- [chdb-go](./tools/chdb/README.md)
