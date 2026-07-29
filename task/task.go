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

package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Id string

func newId() Id {
	return Id(uuid.NewString())
}

type Task struct {
	Id          Id
	Title       string
	Completed   bool
	Priority    Priority
	Scheduled   *time.Time
	Due         *time.Time
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func formatDateField(label string, t *time.Time) string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf(" (%s:%s)", label, t.Format(time.DateOnly))
}

func (t *Task) String() string {
	mark := ' '
	if t.Completed {
		mark = 'x'
	}

	var priority string
	if t.Priority != PriorityNone {
		priority = fmt.Sprintf(" (prio:%s)", t.Priority)
	}

	return fmt.Sprintf(
		"- [%c] %s%s%s%s",
		mark,
		t.Title,
		formatDateField("due", t.Due),
		formatDateField("scheduled", t.Scheduled),
		priority,
	)
}

// Toggle task completion
// 
// Also changes CompletedAt - to now if done, to nil if not done
func (t *Task) Toggle() {
	if t.Completed {
		t.Completed = false
		t.CompletedAt = nil
	} else {
		t.Completed = true
		t.CompletedAt = new(time.Now())
	}
}

// Make a task from a title.
func MakeTask(title string) *Task {
	return &Task{
		newId(),
		title,
		false,
		PriorityNone,
		nil,
		nil,
		time.Now(),
		nil,
	}
}
