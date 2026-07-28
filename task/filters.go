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

package task

import "time"

type Filter func(*Task) bool

// Get a filter checking if a task is scheduled for today
func Today() Filter {
	return func (task *Task) bool {
		if task.Scheduled == nil {
			return false
		}
		y1, m1, d1 := task.Scheduled.Date()
		y2, m2, d2 := time.Now().Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			return true
		}
		return false
	}
}

// Get a filter checking if a task has given completion state
func ByCompleted(completed bool) Filter {
	return func(t *Task) bool {
		return t.Completed == completed
	}
}

// Get a filter checking if a task has given priority
func ByPriority(p Priority) Filter {
	return func(t *Task) bool {
		return t.Priority == p
	}
}

// Get a filter checking if a tasks due date is before given date
// 
// Checks only date (year, month, day)
func DueBefore(when time.Time) Filter {
	return func(t *Task) bool {
		if t.Due == nil {
			return false
		}
		y1, m1, d1 := t.Due.Date()
		y2, m2, d2 := when.Date()
		if y1 < y2 || m1 < m2 || d1 < d2 {
			return true
		}
		return false
	}
}

// Get a filter checking if a task is overdue
func Overdue() Filter {
	return func(t *Task) bool {
		if t.Due == nil {
			return false
		}
		y1, m1, d1 := t.Due.Date()
		y2, m2, d2 := time.Now().Date()
		if y1 < y2 || m1 < m2 || d1 < d2 {
			return true
		}
		return false
	}
}
