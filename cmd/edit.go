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

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit prefix",
	Short: "Edit a task",
	Long:  `Edit a task`,
	Run: func(cmd *cobra.Command, args []string) {
		prefix, e := getPrefix(args)
		if reportError(cmd, e) {
			return
		}

		task, e := tasksList.FindByPrefix(prefix)
		if reportError(cmd, e) {
			return
		}

		oldTask := *task

		// title
		if cmd.Flags().Changed(titleFlagName) {
			if title == "" {
				reportError(cmd, fmt.Errorf("Empty title"))
			} else {
				task.Title = title
			}
		}

		// priority
		if cmd.Flags().Changed("priority") {
			task.Priority = priority
		}

		// scheduled date
		if cmd.Flags().Changed("schedule") {
			scheduledDate, e := parseDateFlag(schedule)
			if reportError(cmd, e) {
				return
			} else {
				task.Scheduled = scheduledDate
			}
		}

		// due date
		if cmd.Flags().Changed("due") {
			dueDate, e := parseDateFlag(due)
			if reportError(cmd, e) {
				return
			} else {
				task.Due = dueDate
			}
		}

		printTask("Was:", &oldTask)
		printTask("Updated task:", task)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	addDueFlag(editCmd, &due)
	addScheduledFlag(editCmd, &schedule)
	addPriorityFlag(editCmd, &priority)
	addTitleFlag(editCmd, &title)
}
