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
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	fmt.Print(convert.ImageFile2ASCIIString("example/images/baozou.jpg", &convert.Options{
		Ratio: 0.5,
	}))
}
```

## example outputs

| Raw Image         | ASCII Image           |
| ------------- |:-------------:| 
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei_ascii.png) |
 ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii.png) |
  ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii.png) |

