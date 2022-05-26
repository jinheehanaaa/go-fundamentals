<h1>**Go Fundamentals for Future Reference**</h1>

**[Things I'm learning now]**
- Minimizing GC Work & Heap Allocation
- Value Sementics 
- Pointer Sementics
- UTF-8 Encoding
- Decoupling
- Polymorphism & Interface
- Inner Type Promotion
- net pkg & temporary interface for error handling
- Wrapping Errors for handling an error
- Data Race
-- Atomic for counter
-- Mutex for block of code (atomically)
- Goroutines
-- IO Bound Workload 
-- CPU Bound Workload with Concurrency & Parallelism

**[Things I do not fully understand yet]**
- Unmarshal
- Decode
- Mocking
- Local Run Queue (LRQ)
- Global Run Queue (GRQ)

**[Questions to ask]**
Q: 
- Why use decoupling?
A: 
- Decoupling separates concrete data & behavior
- We can define same behaviors for each data
- Each data with the same interface will have the common behavior
