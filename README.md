# :foggy:image2ascii

[![Build Status](https://travis-ci.org/qeesung/image2ascii.svg?branch=master)](https://travis-ci.org/qeesung/image2ascii)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/71a3059b49274dde9d81d58cedd80962)](https://app.codacy.com/app/qeesung/image2ascii?utm_source=github.com&utm_medium=referral&utm_content=qeesung/image2ascii&utm_campaign=Badge_Grade_Dashboard)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Image2ASCII is a library that converts images into ASCII images and provides command-line tools for easy use.

![demo](https://github.com/qeesung/image2ascii/blob/master/example/images/lufei.gif?raw=true)

## Installation

```bash
go get github.com/qeesung/image2ascii
```

## CLI usage

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
image2ascii -f example/images/baozou.jpg -w 100 -g 30
```
![demo](https://github.com/qeesung/image2ascii/blob/master/example/images/pikaqiu.gif?raw=true)


convert the image to ascii image by ratio
```bash
# ratio: 0.3
# width: imageWidth * 0.3
# height: imageHeight * 0.3
image2ascii -f example/images/pikaqiu.jpg -r 0.3
```
![demo](https://github.com/qeesung/image2ascii/blob/master/example/images/pikaqiu.gif?raw=true)


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

## Library usage

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

convert options

```golang
type Options struct {
	Ratio          float64 .       // Scale Ratio
	ExpectedWidth  int             // Convert the image with fixed width
	ExpectedHeight int             // Convert the image with fixed height
	FitScreen      bool            // Scale the image to fit the tereminal screen
	Colored bool                   // if convert the image to colored ascii
}
```

supported convert function
```golang
// convert a image object to ascii matrix
func Image2ASCIIMatrix(image image.Image, imageConvertOptions *Options) []string {}

// convert a image object to ascii matrix and then join the matrix to a string
func Image2ASCIIString(image image.Image, options *Options) string {}

// convert a image object by input a string to ascii matrix then join the matrix to a string
func ImageFile2ASCIIString(imageFilename string, option *Options) string {}
```



## Sample outputs

| Raw Image                                                                                       | ASCII Image                                                                                                |
|:-----------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------:|
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei_ascii.png)           |
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei.jpg)      | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/lufei_ascii_colored.png)   |
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)   | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii.png)         |
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu.jpeg)   | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/pikaqiu_ascii_colored.png) |
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)     | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii.png)          |
| ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou.jpg)     | ![](https://raw.githubusercontent.com/qeesung/image2ascii/master/example/images/baozou_ascii_colored.png)  |


## License

This project is under the MIT License. See the [LICENSE](https://github.com/qeesung/image2ascii/blob/master/LICENSE) file for the full license text.
