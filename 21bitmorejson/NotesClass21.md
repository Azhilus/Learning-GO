# Class 21: JSON Operations in GoLang

## Understanding JSON Operations:

1. **Introduction:**
   - JSON (JavaScript Object Notation) is a lightweight data interchange format widely used for data exchange between a server and a web application. This class explores encoding Go structs into JSON and decoding JSON data into Go structs.

2. **Purpose of JSON Operations:**
   - JSON operations are essential for serializing Go data structures into a format that can be easily transmitted or stored, and deserializing JSON data into Go structures for processing.

3. **Key Concepts:**
   - Go's `encoding/json` package provides functions for encoding and decoding JSON. Struct tags are used to customize the encoding and decoding process.

4. **Common Use Cases:**
   - JSON operations are employed in scenarios such as communicating with web APIs, storing configuration data, and exchanging data between different systems.

## JSON Operations in Go:

5. **Encoding JSON:**
   - Use `json.Marshal` to encode Go structs into JSON. Example:
   ```go
   finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
   ```

6. **Decoding JSON:**
   - Use `json.Unmarshal` to decode JSON data into Go structs. Example:
   ```go
   var lcoCourse course
   json.Unmarshal(jsonDataFromWeb, &lcoCourse)
   ```

7. **Checking JSON Validity:**
   - Use `json.Valid` to check the validity of JSON data. Example:
   ```go
   checkValid := json.Valid(jsonDataFromWeb)
   ```

8. **Decoding into a Map:**
   - Decode JSON data into a map for more flexibility. Example:
   ```go
   var myOnlineData map[string]interface{}
   json.Unmarshal(jsonDataFromWeb, &myOnlineData)
   ```

9. **Iterating through JSON Map:**
   - Iterate through the decoded JSON map. Example:
   ```go
   for k, v := range myOnlineData {
       fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
   }
   ```

10. **Conclusion:**
    - JSON operations in Go are fundamental for handling data interchange. Encoding Go structs into JSON and decoding JSON into Go structs enable seamless communication with web services and APIs. Understanding these operations is crucial for developers working with data-driven applications.
