A stack is an ordered data structure that follows the Last-In-First-Out (LIFO) principle. Stacks are most easily implemented in Golang using slices: An element is pushed to the stack with the built-in append function. The element is popped from the stack by slicing off the top element.

![stack](image.png)

stack have tow opration 

1. **push**:- store the data in the stack we will use the push operation to store the data in the stack and it will store the data in the next empty index in the stack 


2. **pop**:- to remove data from the stack we will use pop operation and it will remove the first element from the stack as the stack is working in FIFO mechanism 