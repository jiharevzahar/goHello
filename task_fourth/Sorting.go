package main

import (
	"fmt"
	"sort"
	"time"
	"unicode/utf8"
)

type Person struct {
	firstName string
	lastName  string
	birthday  time.Time
}

type People []Person

func (a People) Len() int {
	return len(a)
}
func (a People) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(a[i].firstName)
	jRune, _ := utf8.DecodeRuneInString(a[j].firstName)

	zRune, _ := utf8.DecodeRuneInString(a[i].lastName)
	xRune, _ := utf8.DecodeRuneInString(a[j].lastName)

	if int32(iRune) == int32(jRune) {
		if int32(zRune) == int32(xRune) {
			return a[i].birthday.Before(a[j].birthday)
		} else {
			return int32(zRune) < int32(xRune)
		}
	} else {
		return int32(iRune) < int32(jRune)
	}
}
func (a People) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func main() {
	groupA := []Person{
		{"Clara", "Jonson", time.Now()}, {"Bndr", "Taiwon", time.Date(2000, 0, 1, 12, 0, 0, 0, time.UTC)}, {"Abc", "cvx", time.Date(1999, 0, 1, 12, 0, 0, 0, time.UTC)},
	}

	sort.Sort(People(groupA))
	fmt.Println(groupA)

	groupB := []Person{
		{"AAA", "CCC", time.Now()}, {"AAA", "BBB", time.Date(2000, 0, 1, 12, 0, 0, 0, time.UTC)}, {"AAA", "AAA", time.Date(1999, 0, 1, 12, 0, 0, 0, time.UTC)},
	}
	sort.Sort(People(groupB))
	fmt.Println(groupB)

	groupC := []Person{
		{"BBB", "BBB", time.Now()}, {"AAA", "AAA", time.Date(2005, 0, 1, 12, 0, 0, 0, time.UTC)}, {"BBB", "BBB", time.Date(1999, 0, 1, 12, 0, 0, 0, time.UTC)},
	}
	sort.Sort(People(groupC))
	fmt.Println(groupC)
}
