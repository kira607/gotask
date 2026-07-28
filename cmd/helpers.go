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
	"time"

	"github.com/kira607/gotask/task"
	"github.com/spf13/cobra"
)

const (
	titleFlagName    = "title"
	scheduleFlagName = "schedule"
	dueFlagName      = "due"
	priorityFalgName = "priority"
)

// tasks collection
var tasksList = &task.Tasks{}

// common flags variables (passed by reference)
var showIds bool
var due string
var schedule string
var title string
var priority task.Priority = task.PriorityNone

func printTask(label string, task *task.Task) {
	if showIds {
		fmt.Printf("%s\n(%s) %s\n", label, task.Id, task)
	} else {
		fmt.Printf("%s\n%s\n", label, task)
	}
}

// Print to stdout given tasks slice
func printTasksList(tasks []*task.Task, showIds bool) {
	for _, task := range tasks {
		if showIds {
			fmt.Printf("(%s) %s\n", task.Id, task)
		} else {
			fmt.Printf("%s\n", task)
		}
	}
}

func printTotal(count int, empty string, label string) {
	if count == 0 {
		fmt.Printf("%s\n", empty)
	} else {
		fmt.Println("---")
		fmt.Printf("%d %s\n", count, label)
	}
}

// Parse a value of a flag that accepts a date.
//
// Understands "today", "tomorrow" and a date in format "YYYY-MM-DD"
//
// If a flag value is "none" date will be nil.
func parseDateFlag(date string) (*time.Time, error) {
	if date == "" {
		return nil, nil
	}

	switch date {
	case "today":
		return new(time.Now()), nil
	case "tomorrow":
		return new(time.Now().Add(time.Hour * 24)), nil
	default:
		t, e := time.Parse(time.DateOnly, date)
		if e != nil {
			return &t, fmt.Errorf("Invalid date format: %s", e)
		} else {
			return &t, nil
		}
	}
}

// Try to report an error.
//
// Prints error cmd help to stdout if error is not nil and returns true.
// If error is nil returns false.
func reportError(cmd *cobra.Command, e error) bool {
	if e != nil {
		fmt.Printf("%s\n\n", e)
		cmd.Help()
		return true
	}
	return false
}

// Get a task id prefix from a command arguments list
//
// Expects args to have only one string - the prerix of a task id
func getPrefix(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("No prefix.")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("Too much arguments. One prefix required.")
	}

	return args[0], nil
}

func addPriorityFlag(cmd *cobra.Command, priority *task.Priority) {
	cmd.Flags().VarP(&priorityFlag{value: priority}, priorityFalgName, "p", fmt.Sprintf("task priority (%s)", task.PrioritiesString()))
}

func addScheduledFlag(cmd *cobra.Command, schedule *string) {
	cmd.Flags().StringVarP(schedule, scheduleFlagName, "s", "", "a scheduled date")
}

func addDueFlag(cmd *cobra.Command, due *string) {
	cmd.Flags().StringVarP(due, dueFlagName, "d", "", "a due date")
}

func addTitleFlag(cmd *cobra.Command, title *string) {
	cmd.Flags().StringVarP(title, titleFlagName, "t", "", "task title")
}
