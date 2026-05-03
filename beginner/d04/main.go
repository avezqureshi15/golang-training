package main

import "fmt"

func stats(grades []float64) (float64, float64,  float64){
	if len(grades) == 0{
		return 0,0,0
	}

	max := grades[0];
	min := grades[0];
	sum := 0.0
	for _,val := range grades{
		if val > max {
			max = val 
		}
		if val < min{
			min = val 
		}
		sum += val 
	}

	avg := sum/ float64((len(grades)))

	return max, min, avg
}

func main() {
	gradebook := map[string][]float64{
		"Avez":  {90, 98, 97},
		"Rahul": {99, 76.95},
		"Neha":  {92, 89, 94},
	}

	var allGrades []float64

	fmt.Println("---------student stats--------")
	for name, grades := range gradebook{
		max, min, avg := stats(grades)
		fmt.Println("Student name ",name)
		fmt.Println("Max Score ",max)
		fmt.Println("Min Score ",min)
		fmt.Println("Avg Score ",avg)

		allGrades = append(allGrades, grades...)
	} 

	fmt.Println(allGrades)

	fmt.Println("_____Class Stats______")
	max,min,avg := stats(allGrades)

	fmt.Println("Class Avg ",avg)
	fmt.Println("Highest Grade ",max)
	fmt.Println("Lowest Grade ",min)
}

