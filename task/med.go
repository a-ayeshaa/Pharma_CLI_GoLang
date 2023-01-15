package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//Seeding
func medSeed() []Medicine {
	med := make([]Medicine, 10)
	for i := 0; i < 10; i++ {
		var m Medicine
		m.id = i
		m.name = "Napa" + strconv.Itoa(i)
		m.price = rand.Intn(1000-500) + 500
		med[i] = m
	}
	return med
}

var medlist = medSeed()

func medDb() []Medicine {
	return medlist
}

type Medicine struct {
	id    int
	name  string
	price int
}

type IMedicine interface {
	getAll() []Medicine
	get(id int) Medicine
	add(med Medicine) Medicine
	delete(id int) bool
	update(med Medicine) Medicine
}

func (medicine Medicine) getAll() []Medicine {
	return medDb()
}

func (medicine Medicine) get(id int) Medicine {
	meds := medDb()

	return meds[id]
}

func (medicine Medicine) add(M Medicine) Medicine {
	medlist = append(medlist, M)
	return M
}

func (medicine Medicine) delete(id int) bool {
	// var med []Medicine
	for i, medval := range medlist {
		if medval.id == id {
			// if id == len(medlist)-1 {
			// 	medlist = append(medlist[:i])
			// 	fmt.Println(medlist, len(medlist))
			// 	return true
			// } else {
				medlist = append(medlist[:i], medlist[i+1:]...)
				// medlist[id]=Medicine {}
				fmt.Println(medlist)
			// }

		}
	}

	return true
}

func (medicine Medicine) update(med Medicine) Medicine {
	for i, medval := range medlist {
		if medval.id == med.id {
			medlist[i] = med
			return medlist[i]
		}
	}
	return med
}
