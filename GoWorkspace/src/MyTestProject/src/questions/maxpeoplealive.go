package questions

import "fmt"

//sample code to run this

// p1 := questions.Person{2000, 2003}
// p2 := questions.Person{2000, 2002}

// pArr := []questions.Person{p1, p2}

// questions.MaxPeopleAlive(pArr)

type Person struct {
	BirthY int
	DeathY int
}

func MaxPeopleAlive(people []Person) {

	m1 := make(map[int]int)
	for _, p := range people {
		GetYearsAddedToTheMap(p, m1)
	}

	max := 0
	maxY := 0

	for k, v := range m1 {
		if v > max {
			max = v
			maxY = k
		}
	}

	fmt.Println(m1)
	fmt.Println(max)
	fmt.Println(maxY)
}

func GetYearsAddedToTheMap(p Person, m1 map[int]int) {
	start := p.BirthY
	end := p.DeathY
	index := start

	for index < end+1 {
		if prev, ok := m1[index]; ok {
			m1[index] = prev + 1
		} else {
			m1[index] = 1
		}
		index = index + 1
	}
}
