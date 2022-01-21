package maxicode

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		desc      string
		expPath   string
		mode      int
		inputData string
	}{
		{
			desc:      "mode 2",
			expPath:   "maxicode_mode_2.png",
			mode:      2,
			inputData: "[)>" + RS + "01" + GS + "96841706672" + GS + "840" + GS + "001" + GS + "1Z12345673" + GS + "UPSN" + GS + "1X2X3X" + GS + "187" + GS + "" + GS + "1/1" + GS + "10" + GS + "N" + GS + "19 SOUTH ST" + GS + "SALTLAKE CITY" + GS + "UT" + RS + EOT,
		},
		{
			desc:      "mode 3",
			expPath:   "maxicode_mode_3.png",
			mode:      3,
			inputData: "[)>" + RS + "01" + GS + "09651147" + GS + "276" + GS + "066" + GS + "1Z12345677" + GS + "UPSN" + GS + "1X2X3X" + GS + "187" + GS + "" + GS + "1/1" + GS + "10" + GS + "N" + GS + "5 WALDSTRASSE" + GS + "COLOGNE" + GS + "" + RS + "" + EOT,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			grid, err := Encode(tc.mode, 0, tc.inputData)
			if err != nil {
				t.Fatalf(err.Error())
			}

			buf := bytes.NewBufferString("")

			err = grid.Draw(35.0).EncodePNG(buf)
			if err != nil {
				t.Fatalf(err.Error())
			}

			expected, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s", tc.expPath))
			if err != nil {
				t.Fatalf("Failed to load testdata picture %q: %v", tc.expPath, err)
			}

			if diff := cmp.Diff(string(expected), buf.String()); diff != "" {
				t.Errorf("Generated label is not equal to expected, diff (-want, +got):\n%s", diff)
			}
		})
	}
}
