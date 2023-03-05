String  - len
        - compare
        - contains
        - substring
        - replace
        - insert
  
Variadic function   - To skip creating a temporary slice just to pass to a func
                    - When the number of input params are unknown
                    - To express your intent to increase the readability

Anonymous function == literal function

An anonymous function is a function which doesn’t contain any name. It is useful when you want to create an inline function.

Closure > A closure is a special type of anonymous function that references variables declared outside of the function itself. In this case we will be using variables that weren’t passed into the function as a parameter, but instead were available when the function was declared.

Defer function > A deferred function func is only executed after its parent func return. Multiple return can be used as well, they run as a stack, one by one

Pointer > Pointer is a variable that stores the memory address of another variable. Pointers have the power to mutate data they are pointing.