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

	"github.com/kira607/gotask/task"
)

type ErrUnknownPriority string

func (e ErrUnknownPriority) Error() string {
	return fmt.Sprintf("Unknown priority: %s", fmt.Sprint(string(e)))
}

// A wrapper around task's Priority with cobra's pflag features
type priorityFlag struct {
	value *task.Priority
}

func (f *priorityFlag) String() string {
	if f.value == nil {
		return ""
	}
	return f.value.String()
}


func (f *priorityFlag) Set(value string) error {
	switch value {
	case "none":
		*f.value = task.PriorityNone
	case "low":
		*f.value = task.PriorityLow
	case "medium":
		*f.value = task.PriorityMedium
	case "high":
		*f.value = task.PriorityHigh
	case "urgent":
		*f.value = task.PriorityUrgent
	default:
		return fmt.Errorf("must be one of: %s (got %q)\n", task.PrioritiesString(), value)
	}
	return nil
}

func (f *priorityFlag) Type() string {
	return "priority"
}
