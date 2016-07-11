> every channel has a block queue
> at first the block queue's type is nil
> then it could change to ch<- and <-ch
> then it could change back to nil

**block and enqueue**:
    the goroutine blocks, and enqueue the goroutine to the block queue on the channel.
**dequeue and unblock**:
    dequeue the block queue, and unblock the poped goroutine. Now both can proceed.

```
A list of goroutines start in order. At first all goroutines are proceeding.
Then one goroutine encounters a chan operation
  if it's being executed on an **unbuffered channel**
    if (or (block queue's type == chan operation)
           (block queue's type == nil))
      block and enqueue.
    else
      dequeue and unblock.

  if it's being executed on a **buffered channel**
    if (or (and (block queue's type == nil)
             (or (buffer is empty and chan operation is <-ch)
                 (buffer is full and chan operation is ch<-)))
           (block queue's type == chan operation))
      block and enqueue.
    else if (and (block queue's type != nil)
                 (block queue's type != chan operation))
      dequeue and unblock.
    else
      proceed
```
