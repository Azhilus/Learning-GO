# Class 19: Working with URLs in GoLang

## Understanding URLs in Go:

1. **Introduction:**
   - Go provides the `net/url` package for parsing, building, and manipulating URLs. Understanding how to work with URLs is crucial for web development and API interactions.

2. **Purpose of URL Operations:**
   - URL operations are essential for extracting information from URLs, building URLs dynamically, and handling query parameters.

3. **Key Concepts:**
   - Go's `url` package allows parsing URLs, accessing their components, and building new URLs.

4. **Common Use Cases:**
   - URL operations are used in scenarios such as parsing query parameters, constructing dynamic URLs, and handling different components of a URL.

## Working with URLs in Go:

5. **Parsing a URL:**
   - Use `url.Parse` to parse a URL and obtain its components. Example:
   ```go
   result, err := url.Parse(myurl)
   ```

6. **Accessing URL Components:**
   - Access various components of the parsed URL, such as scheme, host, path, port, and raw query. Example:
   ```go
   fmt.Println(result.Scheme)
   fmt.Println(result.Host)
   fmt.Println(result.Path)
   fmt.Println(result.Port())
   fmt.Println(result.RawQuery)
   ```

7. **Query Parameters:**
   - Retrieve query parameters using the `Query` method. Example:
   ```go
   qparams := result.Query()
   ```

8. **Building a URL:**
   - Create a new URL using the `&url.URL{}` struct and its components. Example:
   ```go
   partsOfUrl := &url.URL{
       Scheme:   "https",
       Host:     "lco.dev",
       Path:     "/tutcss",
       RawQuery: "user=aman",
   }
   ```

9. **Converting URL to String:**
   - Convert the URL struct to a string using the `String` method. Example:
   ```go
   anotherUrl := partsOfUrl.String()
   ```

10. **Conclusion:**
    - The `net/url` package in Go simplifies working with URLs, offering methods for parsing, accessing components, and building URLs. Understanding these operations is essential for developing robust web applications and interacting with web services.
