package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	fmt.Print(convert.ImageFile2ASCIIString("example/images/baozou.jpg", &convert.Options{
		Ratio: 1,
		FitScreen: true,
		ExpectedHeight:51,
		ExpectedWidth: 185,
	}))
}


// outout
/*
                       .. fLCGGG8G ..
                    . GGGGGGCCCCCCCG008 ..
                 . 0GCCG08@       80GCLLG0G.
               .0GLC0@     @@@@@@@@    80CLG0.
             .0CLG8   @@@@@@@@@@@@@@@@@   @GLC8.
          ..8CL0   @@@@@@@@@1i8@@@GCC@@@@@   Gf0.:
         ; Gf0  @@@@@@@@@@01if@@@@G1ft0@@@@@@@@fL 1
        i LL @@@@@@@@@@@0f1CC@@@8GtL8C1f8@@@@@ @Lf@f
       ; fC@@@@@@@@@@0LftC8GL@@@GtC0CG0L1f0@@@@@@Ct@f
      ,.LC@@@@@@@8GCfLGCf00CL@@@G0008@@@@Gftf8@@@@Ct@t
      .Gt@@@@@@@@Gt1ft;  .tGC@@@0C8GCCf1L0@8C0@@@@@Lf@;
     ;@t8@@@@@@08GLLGCLLfttCL@@@8G8fLGL:,,1GGCC8@@@@fC .
     .Gf@@@@@@@@@@@88000080LC@@@008@8G080GC088@@@@@@@1@
    ;@t0@@@88088@@@@@@@@@@@00@@@GG@@@@888@@@@@88@@@@ LC .
     @t@@@8000008@@@@@@@8CCG@@@@@8@@@@@@@@@@000008@@@8t@1
   ..Gf @@0000000@@@@@8G10@@@@@@@@08@@@@@@@8000000@@@@t8
   ,@fC@@@8000008@@@8GGGtGG8@@@@@@8fL8@@@@@80000008@@@t0..
   ;@fG@@@@@88@@@@@GC0@@8CttL8@8GGGffCG8@@@@880088@@@@fC ,
   ;@fG@@@@@@@@@@@CG8@@@@@@@GLLLLLG8@@0GG8@@@@@@@@@@@@LC ,
   :@fC@@@@@@@@@@8C00@@@8888@@@@@@@@@@@0GC8@@@@@@@@@@@LL ,
   . LL@@@@@@@@@@@0CCftffLLLLLLLLCG8@@8080G@@@@@@@@@@@LC.
    .0f @@@@@@@@@@@8C,C@@@@@@8@@880GLL08@8C@@@@@@@@@@@f0.
     @t@@@@@@@@@@@@@@ftf1iii11tfLL0@1:@@@GG@@@@@@@@@@ t@
     .Cf@@@@@@@@@@@@@@t            :L@@@0G@@@@@@@@@@ 0f ,
     ;@tC@@@@@@@@@@@@@@:.:;;;;:::.,G@@@@8@@@@@@@@@@ 8t@
       @tL@ @@@@GCG@@@@CL0@@@@88Gf8@@@@@@@@@@@@@@@@8t8
        @ff8 @@Ct0Ct8@@LGGCGGGGGG0@@@@8GG8@@@@@@ @Gt8 .
        i GtC8 ff@@Gt@@8LLCCCCCCCC8@@CtLCt0@    0fL8
         ,.@CtC010@@tG@@8GCCGGCCG@@@tf8@@Cf@0GCLL0
            .8C11;8@Gf@@@@@@@@@@@@@Lt@@@01fLLCG0.
               tGiL@0i@@@@@@@@@@@@8i@@@L:;0GC..
             .0L Ct@L;fLCCCGGGGGCL11@@tf0t@
              8f0t001CCLLffftttffCft@fL@CC.
             , CfLftLCCCGGG0088@@@0i8GC0t@
              : 0C10G0088@@@@@@@@@@Ctftt8
                 @t@@@@@@@@@@@@@@@@@@8Gt ;
                . CC@@@@@@@@@@@@@@@@@@@CL ;
 */
