# goavcodec

I just wanna convert WebM to MP4

# Example

```go
package main

import (
    "github.com/otiai10/goavcodec"
)

func main() {

    src, err := os.Open("target.webm")
    if err != nil {
      panic(err)
    }
    defer src.Close()

    dest, err := os.Create("dest.mp4")
    if err != nil {
      panic(err)
    }
    defer dest.Close()

    if err := goavcodec.Convert(src, dest); err != nil {
      panic(err)
    }

}
```
