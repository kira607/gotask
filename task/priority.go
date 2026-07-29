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
	"slices"
	"strings"
)

// A task priority
type Priority int

const (
	PriorityNone = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
	PriorityUrgent
)

// All valid priority values
var AllPriorities []Priority = []Priority{
	PriorityNone,
	PriorityLow,
	PriorityMedium,
	PriorityHigh,
	PriorityUrgent,
}

func (p Priority) String() string {
	switch p {
	case PriorityNone:
		return "none"
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	case PriorityUrgent:
		return "urgent"
	default:
		panic("Unknown priority???")
	}
}

// Check if a priority has a valid value
func (p Priority) IsValid() bool {
	return slices.Contains(AllPriorities, p)
}

func PrioritiesString() string {
	names := make([]string, len(AllPriorities))
	for i, p := range AllPriorities {
		names[i] = p.String()
	}
	return strings.Join(names, ", ")
}
