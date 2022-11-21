package elasticsearch

import "fmt"

func buildQuery(keywords []string, size int, sort bool) map[string]interface{} {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Printf("foods: %v\n", foods)
	return foods
}