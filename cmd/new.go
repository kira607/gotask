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
	"strings"

	"github.com/kira607/gotask/task"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new title",
	Short: "Create a new task",
	Long: `Create a new task.

--due and --scheduled accept date in format YYYY-MM-DD
or literals: 'today', 'tomorrow'
	`,
	Aliases: []string{"add", "n"},
	Run: func(cmd *cobra.Command, args []string) {
		// title
		title := strings.Join(args, " ")
		if title == "" {
			reportError(cmd, fmt.Errorf("Empty title!"))
			return
		}

		// make a newTask
		newTask := task.MakeTask(title)

		// priority
		newTask.Priority = priority

		// scheduled date
		scheduledDate, e := parseDateFlag(schedule)
		if reportError(cmd, e) {
			return
		} else {
			newTask.Scheduled = scheduledDate
		}

		// due date
		dueDate, e := parseDateFlag(due)
		if e != nil {
			fmt.Printf("%s\n\n", e)
			cmd.Help()
			return
		} else {
			newTask.Due = dueDate
		}

		// Add and report
		tasksList.Add(newTask)
		printTask("Added a new task:", newTask)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	addDueFlag(newCmd, &due)
	addScheduledFlag(newCmd, &schedule)
	addPriorityFlag(newCmd, &priority)
}
