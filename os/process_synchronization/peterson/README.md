# Peterson Algorithm


該算法遵守:
- 空閒讓進
- 忙則等待
- 有限等待

但不遵守讓權等待

```
//flag[] is boolean array; and turn is an integer
flag[0]   = false;
flag[1]   = false;
int turn;
```

## Process 0

```
flag[0] = true;
turn = 1;
while (flag[1] == true && turn == 1)
{
    // busy wait
}
// critical section
...
// end of critical section
flag[0] = false;
```

## Process 1

```
flag[1] = true;
turn = 0;
while (flag[0] == true && turn == 0)
{
    // busy wait
}
// critical section
...
// end of critical section
flag[1] = false;
```
