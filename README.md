# Custom Error Handling in Go

A comprehensive Go package for enhanced error handling, providing features like error wrapping, custom error codes, localized error messages, and stack trace capturing.

## Features

- **Custom Error Codes**: Define and use custom error codes for better error handling and client-side error identification.
- **Localized Error Messages**: Support for error messages in multiple languages, including English and Persian.
- **Stack Trace Capturing**: Option to capture and include stack traces in errors for easier debugging.
- **Error Wrapping**: Wrap existing errors with additional context, preserving the original error information.
- **Flexible Configuration**: Configure error properties such as stack trace depth and whether to capture stack traces.

## Installation

To install the package, run the following command in your Go environment:

```sh
go get github.com/mehran-dabi/customerror
```

## Usage

### Creating a New Custom Error

```go
import "github.com/mehran-dabi/customerror"

err := someFunction()
if err != nil {
    customErr := customerror.New(err, customerror.ErrorConfig{
        Code: 1000, // some user-defined error code
        Messages: map[customerror.Lang]string{
            customerror.EnLang: "Resource not found",
            customerror.FaLang: "منبع یافت نشد",
        },
        CaptureStackTrace: true,
    })
    // Use customErr
}
```

### Wrapping an Existing Error

```go
import "github.com/mehran-dabi/customerror"

err := someFunction()
if err != nil {
    wrappedErr := customerror.E(err)
	// Use wrappedErr
}
```

### Contributing

Contributions to the package are always welcome! Please ensure to follow best practices and add tests for new features.

