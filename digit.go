package main

import (
	"fmt"
	"math"
	"flag"
)

func numlength(z int)(int){
	l := 1
	for z >= 10{
		l++
		z = z/10
	}
	return l
}

func addarr(zarr []int)(int){
	tot := 0
	for _,i := range zarr{
		tot += i
	}
	return tot
}

func main(){
	parts := make([]int, 0)

	//Processing Flags
	var NUM int
	flag.IntVar(&NUM, "n", 1, "Input number")
	var GROUP int
	flag.IntVar(&GROUP, "g", 1, "Grouping size")
	var VERB bool
	flag.BoolVar(&VERB, "v", false, "Add a bit more in depth logging")
	flag.Parse()

	//PAD WITH ZEROS
	TOADD := (GROUP)-(numlength(NUM)%GROUP)
	if TOADD == GROUP{TOADD=0}
	NewNum:=int(float64(NUM)*math.Pow(float64(10), float64(TOADD)))

	//Dividing Number into chunks
	for numlength(NewNum) > 1{
		t:=int(math.Pow(float64(10), float64(GROUP)))
		parts = append(parts, NewNum%t)
		NewNum/=t
	}
	if VERB{
		fmt.Println(parts)
	}
	fmt.Println(addarr(parts))
}