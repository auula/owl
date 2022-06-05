// MIT License

// Copyright (c) 2022 Leon Ding <ding@ibyte.me>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package run

import (
	"os"
	"time"

	"github.com/auula/deepscan/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	helpLong = `
 
	Example:
	
	Scan the target data file or directory according to different feature codes 👇
	$ ./deepscan run --dir=/Users/ding/desktop/deepscan/ --md5=86dsxxxxxxf888 --
	out=result.json

	Find a file based on a specific hex value 👇
	$ ./deepscan run --dir=/Users/ding/desktop/tcascan/ --hex=74 63 61 73 63 61 6e 2f
	`
)

var Cmd = cobra.Command{
	Use:   "run",
	Short: "Execute the scanner",
	Long:  color.GreenString(helpLong),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Loading Files...")
		start := time.Now()
		elapsed := time.Since(start)
		log.Info("Scanning time to complete: ", elapsed)
		os.Exit(0)
	},
}
