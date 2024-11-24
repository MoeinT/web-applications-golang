# web-applications-golang
This documentation serves as a beginner-friendly guide to creating a basic web server in Go. It explains key components, how to handle HTTP requests, and how to set up a listener for client connections.

## Hello-world application
### HTTP Handler Function
The handler function processes incoming HTTP requests and sends responses back to the client. Handler functions define what should happen when a user visits a specific URL path. For example, they might return a greeting or dynamically generated content based on the request. An important thing to rememeber is that in order for a function to respond to a request from a web browser, it needs the have the first two parameters below:
1. **http.ResponseWriter:**
    - An interface used to write the HTTP response back to the client.
    - You use methods like Write or Fprintln to send text, HTML, or JSON content as the response.
2. ***http.Request:**
    - A pointer to the http.Request struct that represents the client's request.
    - It contains essential details such as:
        - `Method`: The HTTP method (e.g., GET, POST).
        - `URL.Path`: The requested path (e.g., /, /about).
        - `Query Parameters`: A map of query values (e.g., ?name=John).
        - `Headers`: Information like user-agent or authorization.
        - `Body`: The request body (useful for POST or PUT).
3. **The HTTP Server**
    - A web server is responsible for listening to incoming client requests and routing them to the appropriate handlers. Go’s http.`ListenAndServ` function sets up a server. Here are few important components: 
        - **Purpose**: Starts a web server that listens on a specified port.
        - **Parameters**
            - Port (e.g., :8080): The address and port number to listen on.
            - Handler (e.g., nil): The request multiplexer that routes requests. If nil, Go uses its default multiplexer.
        - **Behavior**:
            - If the server starts successfully, it blocks further execution and continuously waits for incoming requests.
            - If the server fails to start (e.g., due to a port conflict), it returns an error.
### Typical Workflow
- Set Up a Handler: Define how the server responds to specific paths.
- Start the Server: Use http.ListenAndServe to begin listening for requests.
- Access the Server: Use a browser or HTTP client to send requests to the server's URL.
- Process Requests: The handler receives the client's request and responds accordingly.

### Templates
Templates in Go are a way to generate HTML dynamically on the server. The html/template package allows developers to define static HTML with placeholders that can be replaced by dynamic data during runtime. This approach separates the application logic from the presentation layer, making the codebase more maintainable and flexible.

In Go, we can create an HTML file (e.g., home.html) with placeholders for dynamic content. Placeholders are written using ```{{ .FieldName }}```. Here are a few important methods to create a hello-world application using a dynamic template:
1. **template.ParseFiles**
    - Use this function to load the template file in your Go code. It returns a template object and an error object
2. **tmpl.Execute**: 
    - Use this function to execute your template object and return the result back to the client, so it needs object of type `http.ResponseWriter` as an input.
3. **Passing data**
    - The most common and recommended approach is to pass a struct object to the `Execute` function. Specific fields of this struct can be accessed inside the template using the dot notation.

### Template Caching Methodology
We create a template cache as a `map[string]*template.Template` where:
- Each key is the file name of the page template (e.g., home.page.html).
- Each value is a parsed `*template.Template` that combines the page template with any layout templates.
- This approach avoids reparsing templates repeatedly during runtime, improving performance and maintaining separation of concerns between page content and layout structure.
- A key point in this implementation is that we can use the `ParseFiles` method as part of the template package to parse one or multiple html files, but we can also use it an a method of `*template.Template` to append a parsed file to a partially parsed file. We used this concept to parse pages first, and then appending layouts, if any, to them.

### Sharing data between different parts of the application
In Go, two common approaches to sharing data across different parts of an application are global state with update functions and function receivers on structs. 

- A receiver is a special parameter in Go methods that associates the method with a specific type, often a struct, allowing the struct to hold shared data and dependencies. 
- The global state approach involves defining a variable that matches the desired type and providing functions to update or access it, which is simple but risks issues like global state conflicts and poor concurrency safety. 
- In contrast, using function receivers ties methods to a struct that encapsulates the shared data, promoting modularity, testability, and scalability. Function receivers are idiomatic in Go, as they allow for explicit dependencies and better code organization, especially in medium to large applications. Both approaches have use cases, but the choice depends on application complexity and long-term maintainability requirements.