# Temporalis

Go package that provides additional functionality for working with time values beyond what is provided by the standard time package. It includes all of the functions from the time package as well as additional functions for formatting and parsing time values.

## Installation

To install temporalis, use `go get`:

```bash
go get github.com/iamando/temporalis
```

## Usage

To use the module in your Go program, import it using the following code:

```go
import "github.com/iamando/temporalis"
```

`temporalis` can be used just like the standard time package. Here's an example usage of the `temporalis` package:

```go
package main

import (
    "fmt"
    "time"

    "github.com/iamando/temporalis"
)

func main() {
    t := temporalis.Now()
    fmt.Println(temporalis.Format(t, "2006-01-02 15:04:05"))
}
```

In this example, the `temporalis.Now()` function returns the current time, and then the `temporalis.Format` function is used to format that time using a layout string. The resulting formatted string will be in the format `2006-01-02 15:04:05`.

`temporalis` also includes additional functions for formatting and parsing time values. Here's an example usage of the `temporalis.Parse` function:

```go
package main

import (
    "fmt"

    "github.com/your-username/temporalis"
)

func main() {
    str := "2022-05-02 10:30:00"
    t, err := temporalis.Parse("2006-01-02 15:04:05", str)
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return
    }
    fmt.Println(t)
}
```

## Testing

```bash
go test
```

## Support

Temporalis is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

Temporalis is [MIT licensed](LICENSE).
