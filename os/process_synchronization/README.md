# Process Synchronization

## 互斥訪問需要遵守以下條件：

- 空閒讓進: 當 critical section 空閒時，process可以進入
- 忙則等待: 當 critical section 被其他process占用時，process需要等待
- 有限等待: process等待進入 critical section 的時間有限 (保證不會饑渴)
- 讓權等待: process不能進入 critical section 時，應該讓權給其他process