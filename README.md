# maxicode
Library for creating UPS MaxiCodes written in pure Go

## Usage

```

 inputData := "[)>" + RS + "01" + GS + "09651147" + GS + "276" + GS + "066" + GS + "1Z12345677" + GS + "UPSN" + GS + "1X2X3X" + GS + "187" + GS + "" + GS + "1/1" + GS + "10" + GS + "N" + GS + "5 WALDSTRASSE" + GS + "COLOGNE" + GS + "" + RS + "" + EOT
 mode := 3
 eci := 0
 scaleFactor := 35.0
 
 maxicode.Encode(mode, eci, tc.inputData).Draw(scaleFactor).SavePNG("maxicode.png")

```

## Contributors
 
Special thanks for:
* [ZintNET - Barcode Library](https://sourceforge.net/projects/zintnet/). That repository served as a source of inspiration for this particular library