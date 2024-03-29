# Class 5: Time Library in GoLang and Cross-Compilation

## Time Library in GoLang:

1. **Welcome Message:**
   - Introduction to the time package in GoLang, emphasizing its significance in managing time-related data.

2. **Current Time:**
   - The `Now` function from the time package is used to obtain the current time.

3. **Formatting Present Time:**
   - The `Format` method is demonstrated for formatting and presenting various components of the current time. Developers can refer to the documentation for layout string options (e.g., "01-02-2006 15:04 Monday").

4. **Creating Specific Date and Time:**
   - The `Date` function creates a specific date and time, allowing developers to customize each component, including year, month, day, hour, minute, second, nanosecond, and time zone.

5. **Formatting Created Date:**
   - Similar to the present time, the `Format` method is used to format and present specific components of the created date. Developers can tailor the layout string based on their formatting preferences.

6. **Time Package Documentation:**
   - For comprehensive information on the time package, developers should refer to the official Go documentation. It provides detailed explanations of each function, method, and supported layout strings for time formatting.

7. **Modifying Time Format:**
   - Developers can modify the time format by using the `Format` method with different layout strings. Understanding the layout strings is crucial, and the documentation offers a comprehensive guide on the available options.

8. **Layout Strings Reference:**
   - The Go documentation provides a reference guide for layout strings, helping developers create custom time formats for their applications. It includes details on formatting year, month, day, time, and other time-related components.

## Cross-Compilation with GOOS:

9. **Introduction:**
   - GoLang supports cross-compilation, allowing developers to build executables for different operating systems.

10. **GOOS Environment Variable:**
    - The `GOOS` environment variable is used with the `go build` command to specify the target operating system (e.g., "linux", "windows", "darwin").
      ```bash
      GOOS="linux" go build main.go
      GOOS="windows" go build main.go
      GOOS="darwin" go build main.go
      ```

11. **Cross-Compilation Note:**
    - The code includes information about cross-compilation as an additional note. Developers interested in cross-compiling Go programs can explore the related documentation to understand how to build binaries for different operating systems and architectures.

12. **Conclusion:**
    - A solid understanding of the time package and cross-compilation empowers developers to build reliable and cross-platform applications that effectively manage time-related aspects.
