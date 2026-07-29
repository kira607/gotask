/*
Copyright © 2026 kira607 <kirill.lesckin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle prefix",
	Short: "Toggle a task completion",
	Long: `Toggle a task completion by its id. You may first want
to list tasks via any list command (all, overdue, today, etc.) with
an flag for listing ids, and then use an id of a task
(or some first letters of it). For example:

gotask toggle u75be
	`,
	Run: func(cmd *cobra.Command, args []string) {
		prefix, e := getPrefix(args)
		if reportError(cmd, e) {
			return
		}

		task, e := tasksList.FindByPrefix(prefix)
		if e != nil {
			fmt.Printf("Error: %s\n", e)
			return
		}

		task.Toggle()

		printTask("Task toggled:", task)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
