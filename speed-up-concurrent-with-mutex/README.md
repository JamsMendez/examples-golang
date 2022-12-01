We’ll discuss how the use of RWMutex over Mutex can greatly enhance the performance.

Introduction
Multiple threads accessing the same memory at the same time is not desirable.
In Golang we can have several different goroutines that all possibly have access
to the same memory variables, which can lead to race condition.
Mutex, mutual exclusion, along with wait groups to avoid race conditions.
Mutex and RWmutex of golang can be used.

The Mutex in the first is eight bytes. RWMutex is twenty-four bytes.
From official docs : A RWMutex is a reader/writer mutual exclusion lock.
The lock can be held by an arbitrary number of readers or a single writer.
The zero value for a RWMutex is an unlocked mutex.

held if there are pending writers.
semaphore for writers to wait for completing.
number of departing readers.

In a nutshell, readers are not required to wait for one another.
They merely need to wait for the writers who are holding the lock.
Since a just reading function does not change the file contents,
it is acceptable to allow many readers to read the same file at
the same time to enhance the program’s performance.
But writing alters the file’s content, mutually exclusive access
is required; else, overly large mistakes would occur.
A sync.RWMutex is thus suitable for largely read data, and the
resource saved over a sync.Mutex is time.

Every write operation guarded by an RWMutex is O(readers).

Code was ran 10 times on a windows system with intel i5 chip and
the performance was as the average value to complete the concurrent
read-write operation.

The Mutex in the first is eight bytes. RWMutex is twenty-four bytes. 
From official docs : A RWMutex is a reader/writer mutual exclusion lock.
The lock can be held by an arbitrary number of readers or a single writer.
The zero value for a RWMutex is an unlocked mutex.

RWmutex should be used with wait groups when we have a single writer,
to alter shared memory, and multiple reader.
