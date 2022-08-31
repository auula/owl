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

package search

import (
	"fmt"
	"os"

	"github.com/auula/owl/scan"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	helpLong = `
 
	Example:
	
	Scanning and searching of signature files 👇
	$ ./owl search --code=xxxxxxxxx
	`
)

var path, code string

var Cmd = cobra.Command{
	Use:   "search",
	Short: "Searching of signature files",
	Long:  color.GreenString(helpLong),
	Run: func(cmd *cobra.Command, args []string) {
		code = os.Getenv("FEATURE_CODE")
		path = os.Getenv("SOURCE_DIR")
		scan.Exec(func() error {
			scanner := &scan.Scanner{
				Path:    path,
				Code:    code,
				Matcher: new(scan.Md5Matcher),
			}
			restca := make([]scan.ResultElement, 10)
			if res, err := scanner.Search(code); err != nil {
				return err
			} else {
				for _, v := range res {
					restca = append(restca, scan.ResultElement{
						Path:   v.Path,
						Line:   "0",
						Column: "0",
						Msg:    fmt.Sprintf("匹配文件路径为：%v", v.Path),
						Rule:   "",
						Refs:   []scan.Ref{},
					})
				}
				scan.SaveFile("result.json", scanner, restca)
			}
			return nil
		})
	},
}

func init() {
	Cmd.Flags().StringVar(&code, "code", "", "Requires search signature")
	Cmd.Flags().StringVar(&path, "path", "", "The file path where the md5 value needs to be obtained")
}
