# Class 6: Memory Management in GoLang

## Introduction
In this class, we explore the memory management mechanisms in GoLang, including the initialization of memory using `new()` and `make()`. Additionally, we delve into the Garbage Collector (GC) and the related `GOGC` variables, along with the `NumCPU` function.

https://pkg.go.dev/runtime

## Memory Management Overview

### 1. Memory Initialization with `new()` and `make()`
- **`new()` Function:**
  - Used for creating a new instance of a value type or a pointer to a value type.
  - Allocates memory and returns a pointer to the zero-initialized value.
  - Suitable for basic types like integers, floats, and user-defined types.
  - Zeroed Storage 

- **`make()` Function:**
  - Primarily used for initializing slices, maps, and channels.
  - Allocates and initializes memory, returning an initialized instance of the specified composite type.
  - Essential for complex data structures that require initializations, such as dynamic arrays or concurrent communication channels.
  - Non-Zeroed Storage

### 2. Storage Modes in `make()` Function
- **Slices:**
  - Internally uses an array to store elements.
  - Dynamically resizes the array as needed.

- **Maps:**
  - Utilizes a hash table to store key-value pairs.
  - Dynamically adjusts the size of the hash table.

- **Channels:**
  - Employs a queue-like structure for communication between goroutines.
  - Supports buffered and unbuffered channel types.

## Garbage Collector (GC) in GoLang

### 1. Automatic Memory Management
- GoLang features automatic memory management, meaning developers don't manually allocate or deallocate memory.

### 2. Garbage Collector (GC)
- Responsible for identifying and collecting unused memory (garbage).
- Prevents memory leaks and ensures efficient use of resources.

### 3. `GOGC` Environment Variable
- Governs the percentage of heap that triggers garbage collection.
- Default value is typically set to 100, meaning garbage collection is triggered when the heap is 100% full.
- Setting GOGC = off disables GC entirely

### 4. `NumCPU` Function
- Returns the number of logical CPUs available for the current machine.
- The set of available CPUs is checked by querying the OS at process startup.
- Changes to OS CPU allocation after process startup are not reflected
- Useful for optimizing the performance of concurrent programs by allocating appropriate resources.