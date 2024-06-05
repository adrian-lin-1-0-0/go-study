# Semaphore

## Binary Semaphores

```c
int S = 1;

void wait() {
    while (S <= 0);
    S--;
}

void signal() {
    S++;
}

```

## Counting Semaphores

```c
typedef struct {
    int value;
    struct process *list;
} semaphore;

void wait(semaphore *s) {
    s->value--;
    if (s->value < 0) {
        block(s->list);
    }
}

void signal(semaphore *s) {
    s->value++;
    if (s->value <= 0) {
        wakeup(s->list);
    }
}

```