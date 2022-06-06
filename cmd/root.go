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

package cmd

import (
	"fmt"
	"os"

	"github.com/auula/woodpecker/cmd/hex"
	"github.com/auula/woodpecker/cmd/md5"
	"github.com/auula/woodpecker/cmd/run"
	"github.com/auula/woodpecker/cmd/version"
	"github.com/auula/woodpecker/scan"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	bannerStr = `
	 _    _  _____  _____  ____  ____  ____  ___  _  _  ____  ____ 
	( \/\/ )(  _  )(  _  )(  _ \(  _ \( ___)/ __)( )/ )( ___)(  _ \
	 )    (  )(_)(  )(_)(  )(_) ))___/ )__)( (__  )  (  )__)  )   /
	(__/\__)(_____)(_____)(____/(__)  (____)\___)(_)\_)(____)(_)\_)
         A dependency module feature scanning detection tool.
`
)

var banner string = color.BlueString(bannerStr)

var rootCmd = &cobra.Command{
	Use:   "deepscan",
	Short: "deepscan is a dependency module feature scanning detection tool.",
	Long:  banner,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(&version.Cmd)
	rootCmd.AddCommand(&run.Cmd)
	rootCmd.AddCommand(&md5.Cmd)
	rootCmd.AddCommand(&hex.Cmd)
}

var CommonTables = func() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"🔢", "📃File", "🧬MD5"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor})
	return table
}

func WriteTables(table *tablewriter.Table, res []*scan.Result) {
	var list [][]string
	for i, v := range res {
		list = append(list, []string{
			fmt.Sprintf("%d", i+1), color.GreenString(v.Path), color.RedString(v.Code),
		})
	}
	table.AppendBulk(list)
	table.Render()
}
