# image2ascii

Convert image to ASCII

## installation

```bash
go get https://github.com/qeesung/image2ascii
```

## usage

```golang
package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("example/images/lufei.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	asciiImage := convert.Image2ASCIIString(img, &convert.Options{
		Ratio: 0.2,
	})
	fmt.Print(asciiImage)
}
```

## example outputs

| Raw Image         | ASCII Image           |
| ------------- |:-------------:| 
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei_ascii.png) |
 ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii.png) |
  ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii.png) |

