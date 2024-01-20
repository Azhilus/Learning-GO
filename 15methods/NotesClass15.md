# Class 15: Methods in GoLang

## Understanding Methods:

1. **Introduction:**
   - Methods in Go provide a way to associate functions with specific types, enhancing the object-oriented capabilities of the language.

2. **Purpose of Methods:**
   - Methods enable the definition of behavior associated with a type, promoting code organization, encapsulation, and readability.

3. **Key Concepts:**
   - Go supports both value receiver and pointer receiver methods, allowing for methods to modify the state of the receiver.

4. **Common Use Cases:**
   - Methods are employed to define behavior specific to a type, facilitating a more intuitive and modular code structure.

## Methods in Go:

5. **Method Declaration:**
   - Declare a method using the `func` keyword followed by the receiver type, the method name, parameters, return type (if any), and the method body. Example:
   ```go
   type Circle struct {
       Radius float64
   }

   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }
   ```

6. **Value Receiver Methods:**
   - Value receiver methods operate on a copy of the struct and do not modify the original instance. Example:
   ```go
   func (p Point) Move(x, y float64) Point {
       p.X += x
       p.Y += y
       return p
   }
   ```

7. **Pointer Receiver Methods:**
   - Pointer receiver methods operate on the actual instance and can modify its state. Example:
   ```go
   func (c *Counter) Increment() {
       c.Value++
   }
   ```

8. **Promoting Methods:**
   - Methods with pointer receivers can be called on both pointer and value instances of a type. Go automatically converts the value receiver to a pointer receiver when necessary.

9. **Method Sets:**
   - A method set is the set of methods that can be called on a type. It includes both value and pointer receiver methods.

## Traits and Functions:

10. **Choosing Between Value and Pointer Receivers:**
    - Decide between value and pointer receivers based on whether the method needs to modify the state of the receiver.

11. **Method Overloading in Go:**
    - Go does not support traditional method overloading. However, methods with the same name can be defined on different types.

12. **Method Composition:**
    - Methods facilitate composition by allowing types to embed other types and inherit their methods.

13. **Effective Use of Methods:**
    - Utilize methods to encapsulate behavior within types, improving code organization and readability.

14. **Conclusion:**
    - Methods in Go provide a powerful mechanism for associating behavior with types. Understanding value and pointer receiver methods, method sets, and their implications is crucial for designing effective and modular Go programs.
