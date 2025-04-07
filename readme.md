## Go channel patterns:
- `Barrier`, wait for all workers to finish on the breakpoint
- `Done`, wait for a task to be shut down. In case if it's needed just to send cancellation signal it's better to use contexts.
- `DoneWithStruct`, the same as `Done` but the work with the channel is encapsulated
- `ErrGroup` - the simple implementation of the error group, which waits for all tasks to complete unless there is an error.
- `FanIn` - merge multiple channels into one
- `FanOut` - split the work from one channel between multiple channels
- `Filter` - discard some values from the channel based on the predicate
- `First` - Batch processes and wait the first result, other tasks results are ignored.
- `Future` - implementation of the Future pattern, run a task and wait for the result somewhere
- `LeakyBucket` - rate limiting algorithm
- `OrChannel` - like select, but with dynamic number of arms
- `OrDone` - read from the channel until the done signal is received
- `Pipe` - chain of transformations of channels
- `Promise` - then / catch implementation with chaining
- `Shutdown` - listen of os signals to graceful shutdown
- `SingleFlight` - in case if there are a lot of requests to the same resource, it's better to pass one worker and wait for the result in other workers
- `Tee` - like `FanOut`, but send the same value to multiple channels
- `Transformer` - transform the values in the channel

## Semaphore implementations:
- Atomics + Cond - the most efficient
- Mutex + Cond
- Channel - the simplest to implement