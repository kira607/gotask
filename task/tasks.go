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
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// In-memory Tasks collection.
type Tasks struct {
	tasks map[Id]*Task
}

// Create a new task with a given title
func (t *Tasks) New(title string) *Task {
	task := MakeTask(title)
	t.tasks[task.Id] = task
	return task
}

// Add a new task
func (t *Tasks) Add(task *Task) {
	t.tasks[task.Id] = task
}

// Remove a task by id
func (t *Tasks) Remove(id Id) {
	delete(t.tasks, id)
}

// Update a task by id
func (t *Tasks) Update(id Id, updated *Task) {
	updated.Id = id
	t.tasks[id] = updated
}

// Get all tasks
func (t *Tasks) All() []*Task {
	return t.Filter()
}

// Get a slice of tasks each of which match all filters given.
func (t *Tasks) Filter(fs ...Filter) []*Task {
	var result []*Task

	for _, task := range t.tasks {
		if matchesAll(task, fs) {
			result = append(result, task)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.Before(result[j].CreatedAt)
	})
	return result
}

func matchesAll(task *Task, filters []Filter) bool {
	for _, f := range filters {
		if !f(task) {
			return false
		}
	}
	return true
}

// Get a slice of tasks scheduled for today
func (t *Tasks) Today() []*Task {
	return t.Filter(Today())
}

// Get a slice of overdue tasks
func (t *Tasks) Overdue() []*Task {
	return t.Filter(Overdue())
}

// Find a task by Id prefix
func (t *Tasks) FindByPrefix(prefix string) (*Task, error) {
	var matches []*Task
	for id, task := range t.tasks {
		if strings.HasPrefix(string(id), prefix) {
			matches = append(matches, task)
		}
	}
	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("no task matching %q", prefix)
	case 1:
		return matches[0], nil
	default:
		return nil, fmt.Errorf("ambiguous prefix %q: %d matches", prefix, len(matches))
	}
}

func (t *Tasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.tasks)
}

func (t *Tasks) UnmarshalJSON(data []byte) error {
	var tasks map[Id]*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}
	t.tasks = tasks
	return nil
}
