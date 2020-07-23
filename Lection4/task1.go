package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].birthDay == p[j].birthDay {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}

		return p[i].firstName < p[j].firstName
	}

	return p[j].birthDay.Unix() < p[i].birthDay.Unix()
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	const shortForm = "2006-Jan-02"
	ivanIvanovDate, _ := time.Parse(shortForm, "2005-Aug-10")
	ivanIvanovDateAlt, _ := time.Parse(shortForm, "2003-Aug-05")
	artiomIvanovDate, _ := time.Parse(shortForm, "2005-Aug-10")
	
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Ivan", "Ivanov", ivanIvanovDateAlt},
		{"Artiom", "Ivanov", artiomIvanovDate},
	}
	sort.Sort(p)
	for i:=range(p){
		fmt.Printf("%s %s | %s \n", p[i].firstName,p[i].lastName, p[i].birthDay.Format(shortForm))
	}
}
