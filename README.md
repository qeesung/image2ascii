# :foggy:image2ascii

Convert image to ASCII

## installation

```bash
go get https://github.com/qeesung/image2ascii
```

## cli usage

```bash
image2ascii version: image2ascii/1.0.0
>> HomePage: https://github.com/qeesung/image2ascii
>> Issue   : https://github.com/qeesung/image2ascii/issues
>> Author  : qeesung
Usage: image2ascii [-s] -f <filename> -r <ratio> -w <width> -g <height>

Options:
  -c    Colored the ascii when output to the terminal (default true)
  -f string
        Image filename to be convert
  -g int
        Expected image height, -1 for image default height (default -1)
  -r float
        Ratio to scale the image, ignored when use -w or -g (default 1)
  -s    Fit the terminal screen, ignored when use -w, -g, -r (default true)
  -w int
        Expected image width, -1 for image default width (default -1)
```

convert the image to ascii image with fixed width and height

```bash
# width: 100
# height: 30
image2ascii -f example/images/lufei.jpg -w 100 -g 30
```

convert the image to ascii image by ratio
```bash
# ratio: 0.3
# width: imageWidth * 0.3
# height: imageHeight * 0.3
image2ascii -f example/images/lufei.jpg -r 0.3
```

convert the image to ascii fit the screen
```bash
image2ascii -f example/images/lufei.jpg -s
```

convert the image without the color
```bash
image2ascii -f example/images/lufei.jpg -s -c=false
```

convert the image disable fit the screen
```bash
image2ascii -f example/images/lufei.jpg -s=false
```

## library usage

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
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei_ascii_colored.png) |
 ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii.png) |
 ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii_colored.png) |
  ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii.png) |
  ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii_colored.png) |

