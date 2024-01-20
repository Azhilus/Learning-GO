# Class 11: Structs in GoLang

## Understanding Structs:

1. **Introduction:**
   - Structs in Go are composite data types that allow grouping different types of variables under a single name. They provide a way to create custom data structures.

2. **Purpose of Structs:**
   - Structs are used to model real-world entities by combining variables of different types into a single unit, enabling a more organized representation of data.

3. **Key Concepts:**
   - Struct fields can be of any data type, including other structs, creating a flexible and extensible way to structure data.

4. **Common Use Cases:**
   - Structs are suitable for representing entities with multiple attributes, such as a person with a name, age, and address.

## Struct Operations in Go:

5. **Declaration:**
   - Declare a struct using the `type` keyword and define its fields. Example: `type Person struct { Name string; Age int }`.

6. **Initialization:**
   - Initialize a struct with values for its fields. Example: `person := Person{Name: "Alice", Age: 25}`.

7. **Accessing Fields:**
   - Access struct fields using the dot notation. Example: `name := person.Name`.

8. **Updating Fields:**
   - Update struct fields by assigning new values. Example: `person.Age = 26`.

9. **Nested Structs:**
   - Create structs within structs to model complex relationships. Example: `type Address struct { City string; PostalCode string }` and `type Person struct { Name string; Age int; Address Address }`.

10. **Anonymous Structs:**
    - Create anonymous structs on the fly without defining a type. Example: `user := struct { Name string; Age int }{Name: "Bob", Age: 30}`.

11. **Comparing Structs:**
    - Structs are comparable if their fields are comparable. Example: `if person1 == person2 { /* ... */ }`.

## Traits and Functions:

12. **Custom Methods:**
    - Define custom methods for structs to add behavior. Example: 
    ```go
    func (p *Person) SayHello() {
        fmt.Println("Hello, my name is", p.Name)
    }
    ```

13. **Modularity:**
    - Structs enhance modularity by encapsulating related data and behavior in a single unit.

14. **Data Organization:**
    - Structs improve the organization of data, making it easier to manage and understand.

15. **Effective Use of Structs:**
    - Utilize structs to represent entities, model relationships, and encapsulate behavior, fostering clean and modular code.

16. **Conclusion:**
    - Structs in Go are a fundamental building block for creating custom data types. Understanding their declaration, initialization, and usage, along with incorporating custom methods, is essential for building well-structured and maintainable programs.
