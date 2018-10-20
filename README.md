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
	"github.com/qeesung/image2asicc/convert"
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

output
```text
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@888008888@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@8GCCCCCCCCLCCCCCG8@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@8GCCLLLLLLLLLLLLLLLLLCG8@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@0CCLLLffLLLLLLLLLffffffLLLC8@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@0LLffffffttffffffffffLftfLfftG@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@0L@@@8GLfffffLfffLLLLCffLLfffftfffftfL8@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@;iG0CfCfLCLffffffffffffffffffffLfLffff8@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@0GL  ;tLLfttfLLLLLLfffffffftftttttftfLLtfG8@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@80CLfCt.  ,:1ti;;itttt1111ii111tfffft111ttfffLLG8@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@8LLfffLLLf,.   .,,.   . ,ti ., ...,:;i1ttttttttftffC8@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@Gfffffft1iii. ..      . ,fLt..1;  ..   ..:ittttttttfLfG@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@8LLfffti:,.       . . ,:  tGCC; ;L; .;, .    ,itttttttfLtf@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@0LftfLfffft1, .       .f: :CCCCf,.LL: ,t, . .   ,1t1111ttttL@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@0LftLfti:,.        ..  tf..tGCCCLt iGf, tL,  .... .1tt11tttfL0@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@Lftfffffti;:.     .,  iG:;;LGCGCCL1.LG1 ;Ct, ..  . .:1tttttfLC@@@@@@@@@@@@@@@@@
@@@@@@@@@@@8ffffffffLLt;, .   :, .f1,tiCCCCCGCLi;Cf:,ff1..: .   .,:1ttttLf8@@@@@@@@@@@@@@@@
@@@@@@@@@@@8ffttffLfi,.,.     ;, 1f:ftifCCCCCGLC;;Ct:CGC; ;.    .i1ittttt1@@@@@@@@@@@@@@@@@
@@@@@@@@@@@8ft1tfft1i1ti      1..CffCCfftCGCCCGfftLC1CCGf.1,     .tt111tif@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@Ltttttttt11..,    f:;CLCLLLCCCGCCCGCCCfCCCCCL;t,    . :tt1t1t8@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@811tttt11t; i:  . ft1088@8GLCGCCGGCGCC8@@@@8Gtt,.   .:.i1t1f8@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@0t1tttttt;;f,,. ,1C0@@@88@8LCGCCCGCG@@8L0@@@0ti.  . ;;it1L@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@L111ttt1t1:i. ;tG@@@@;.G@8CCCCCCC8@@C.f@@@8t1.., ;:11L0@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@8Cfttttt11t: ,GG@@@@0G@@8CCCCCCC0@@@@@@@@L0GCff11t0@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@80CffffLff,LGG8@@@@@0GCGCGfCGCC08@@80Ct80CG1L1itL0@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@880GGLf1ifLff:10CLCG0GCCCCCC01;CGCCLLfLfLLCCCGffi::;i1C0@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@GCCCCL1iiiifGCC1:GfGCCCCCGGCCCCt1CCCCCCLCCffCGGLt1;;;i1iiitfG8@@@@@@@@@@@@@@@
@@@@@@@@@@@@8ftLCf1iiii;:;CGLt;LfG8000GCCCCCCCGCCCCGG08GC1CGCf1;;;i;;;;iiiiif@@@@@@@@@@@@@@
@@@@@@@@@@@0LC880GGCf11ii;iCLt,;CLG0G8088888000888888G0CfLGCff;;;i;::i1ttfftiL@@@@@@@@@@@@@
@@@@@@@@@0CG80GCCGCGGLf1i1;;ffftfGCG00088888888888800GGfC0GLf1:;;;;:1tfttttLLL0@@@@@@@@@@@@
@@@@@@@8CL08GCCGGGCCCCLfii1;;ii;;1CGCCGGGGGGGGGGGGCCCLtL0CGLt:::,::i11111tfLCGCC08@@@@@@@@@
@@@@@80GG0GGGGGGGGGCLLLf11fffL1::::1LCGGCCCCCCCCCGGGL00GCCCCL111fLCCCCCCCCCCCCGCCC8@@@@@@@@
@@@0GGGGGCCCLLfLLLCLCGGCLLLftfCLti::t11fCGGGGGGGGCf;G8CCGCCCGGGGGGG00CLLCCCfCGGGGGG008@@@@@
@GCGGGCCCLCCCCGCCLL1fCCCLLffCLffLCLt1tti::itfff1;, :8GCGCCGCCCCCCCCCGCf1ftt1tfLCCCCGGG08@@@
GGGGCCCG00800000000000GCGCfffCCCLCCGCftffi:. .    .t@GCCCCCGCCLLffttffttfffCGGG00GCLLGGG8@@
0GCG000GGGCCCCCCCCCGGGGGGGGGGCCGCLLfLCL11ff1;,,,,;tG8GCCCCCCCLLLLLLLfLt1L00000GG000GCfLGG0@
CCCGGCCCCCGGGGGGGGCGGGCCCCCGGGGGGG000000GG0000000880CCCCGCCCGGGGGGGGGGCLCLCGCGGGCCGG0GLLGG0
CCCCCCGCGCGGGGGCGCCGGCCCCGCCCCCCCCCCCGGGGGGGGG0GGGCCGGGGGGGGGGGGGGGGGGGfLffGGGGGGGGGGGGGGGG
```