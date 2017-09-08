# goavcodec

I just wanna convert WebM to MP4

# Example Code

```go
package main

import (
    "fmt"

    "github.com/otiai10/goavcodec/v0/goavcodec"
)

func main() {
    err := goavcodec.Convert("src.webm", "dest.mp4");
    fmt.Println(err)
}
```
// TODO: [v0/goavcodec](https://github.com/otiai10/goavcodec/tree/master/v0) package is just a command wrapper. It should be replaced by Cgo in the future.

# Example Application

- [webm2mp4](https://github.com/otiai10/webm2mp4) to convert WebM to MP4
  - see now on https://webm2mp4.herokuapp.com/
