# Class 2: Web Request Verbs in GoLang

## Understanding Web Request Verbs:

1. **Introduction:**
   - Web request verbs, such as GET, POST, and POST Form, are fundamental to interacting with web services and APIs. This class explores examples of making various types of HTTP requests in Go.

2. **Purpose of Web Request Verbs:**
   - Different HTTP request verbs are used for specific operations, such as fetching data with GET, sending JSON data with POST, and submitting form data with POST Form.

3. **Key Concepts:**
   - Go's `net/http` package provides functions like `http.Get`, `http.Post`, and `http.PostForm` for making different types of HTTP requests.

4. **Common Use Cases:**
   - Web request verbs are employed in scenarios such as fetching data, sending JSON payloads, and submitting form data to web servers.

## Web Request Verbs in Go:

5. **GET Request:**
   - Use `http.Get` to perform a simple GET request. Example:
   ```go
   response, err := http.Get("http://localhost:8000/get")
   ```

6. **Handling GET Response:**
   - Read and handle the response from a GET request. Example:
   ```go
   content, _ := ioutil.ReadAll(response.Body)
   ```

7. **POST JSON Request:**
   - Use `http.Post` for sending JSON data in a POST request. Example:
   ```go
   requestBody := strings.NewReader(`
       {
           "coursename": "Let's go with golang",
           "price": "0",
           "platform": "learnCodeOnline.in"
       }
   `)
   response, err := http.Post("http://localhost:8000/post", "application/json", requestBody)
   ```

8. **POST JSON Response Handling:**
   - Read and handle the response from a POST JSON request. Example:
   ```go
   content, _ := ioutil.ReadAll(response.Body)
   ```

9. **POST Form Request:**
   - Use `http.PostForm` for submitting form data in a POST request. Example:
   ```go
   data := url.Values{}
   data.Add("firstname", "hitesh")
   data.Add("lastname", "choudhary")
   data.Add("email", "hitesh@go.dev")
   response, err := http.PostForm("http://localhost:8000/postform", data)
   ```

10. **POST Form Response Handling:**
    - Read and handle the response from a POST Form request. Example:
    ```go
    content, _ := ioutil.ReadAll(response.Body)
    ```

11. **Conclusion:**
    - Understanding different web request verbs in Go is crucial for building applications that interact with web services and APIs. The provided code examples demonstrate how to perform GET, POST JSON, and POST Form requests, along with handling their responses.