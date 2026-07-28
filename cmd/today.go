/*
Copyright © 2026 kira607 <kirill.lesckin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/kira607/gotask/task"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show tasks scheduled for today",
	Run: func(cmd *cobra.Command, args []string) {
		s, e := cmd.Flags().GetBool("show-ids")
		if e != nil {
			panic(fmt.Sprintf("%s", e))
		}

		d, e := cmd.Flags().GetBool("done")
		if e != nil {
			panic(fmt.Sprintf("%s\n", e))
		}

		filters := []task.Filter{task.Today()}
		if !d {
			filters = append(filters, task.ByCompleted(false))
		}

		tasks := tasksList.Filter(filters...)

		printTasksList(tasks, s)
		printTotal(len(tasks), "No tasks scheduled for today", "Task(s)")
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().BoolP("done", "d", false, "Also show done tasks")
}
