package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}

	for {
		fmt.Println(days[i])
		i++
		if i == len(days) {
			break
		}
	}

	for i := 0; i < len(days); i++ {
		if i == 3 {
			continue
		}
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		if i == 3 {
			break
		}
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		if i == 3 {
			goto mylabel
		}
		fmt.Println(days[i])
	}
mylabel:
	fmt.Println("End of loop")

	for i := 0; i < len(days); i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(days[i], j)
		}
	}

	for i := 0; i < len(days); i++ {
		for j := 0; j < 2; j++ {
			if i == 2 {
				break
			}
			fmt.Println(days[i], j)
		}
	}

	for i := 0; i < len(days); i++ {
		for j := 0; j < 2; j++ {
			if i == 2 {
				continue
			}
			fmt.Println(days[i], j)
		}
	}

}
