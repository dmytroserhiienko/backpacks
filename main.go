package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Item struct {
	name   string
	weight int
	value  int
}

type Individual struct {
	bits []int
}

const (
	MAX_KNAPSACK_WEIGHT = 400
	MUTATION_RATE       = 0.1
	CROSSOVER_RATE      = 0.8
	REPRODUCTION_RATE   = 0.2
)

var (
	items = []Item{
		{"Map", 9, 150},
		{"Compass", 13, 35},
		{"Water", 153, 200},
		{"Sandwich", 50, 160},
		{"Glucose", 15, 60},
		{"Cup", 68, 45},
		{"Banana", 27, 60},
		{"Apple", 39, 40},
		{"Cheese", 23, 30},
		{"Beer", 52, 10},
		{"Sunscreen", 11, 70},
		{"Camera", 32, 30},
		{"T-shirt", 24, 15},
		{"Pants", 48, 10},
		{"Umbrella", 73, 40},
		{"Waterproof pants", 42, 70},
		{"Waterproof coat", 43, 75},
		{"Wallet", 22, 80},
		{"Sunglasses", 7, 20},
		{"Towel", 18, 12},
		{"Socks", 4, 50},
		{"Book", 30, 10},
	}
)

func generateInitialPopulation(count int) []Individual {
	population := make([]Individual, 0)

	for len(population) < count {
		bits := make([]int, len(items))
		for i := range bits {
			bits[i] = rand.Intn(1)
		}
		population = append(population, Individual{bits})
	}

	return population
}

func (ind *Individual) fitness() int {
	totalValue := 0
	totalWeight := 0

	for i, bit := range ind.bits {
		if bit == 1 {
			totalValue += items[i].value
			totalWeight += items[i].weight
		}
	}

	if totalWeight <= MAX_KNAPSACK_WEIGHT {
		return totalValue
	}

	return 0
}

func selection(population []Individual) []Individual {
	parents := make([]Individual, 2)

	for i := 0; i < 2; i++ {
		randIndex1 := rand.Intn(len(population))
		randIndex2 := rand.Intn(len(population))
		if population[randIndex1].fitness() > population[randIndex2].fitness() {
			parents[i] = population[randIndex1]
		} else {
			parents[i] = population[randIndex2]
		}
	}

	return parents
}

func crossover(parents []Individual) []Individual {
	n := len(parents[0].bits)
	crossoverPoint := rand.Intn(n)

	child1 := append(parents[0].bits[:crossoverPoint], parents[1].bits[crossoverPoint:]...)
	child2 := append(parents[1].bits[:crossoverPoint], parents[0].bits[crossoverPoint:]...)

	return []Individual{{child1}, {child2}}
}

func mutate(individuals []Individual) {
	for i := range individuals {
		for j := range individuals[i].bits {
			if rand.Float64() < MUTATION_RATE {
				individuals[i].bits[j] = 1 - individuals[i].bits[j]
			}
		}
	}
}

func nextGeneration(population []Individual) []Individual {
	nextGen := make([]Individual, 0)

	for len(nextGen) < len(population) {
		children := make([]Individual, 0)

		parents := selection(population)

		if rand.Float64() < REPRODUCTION_RATE {
			children = parents
		} else {
			if rand.Float64() < CROSSOVER_RATE {
				children = crossover(parents)
			}
			mutate(children)
		}

		nextGen = append(nextGen, children...)
	}

	return nextGen[:len(population)]
}

func solveKnapsack() Individual {
	rand.Seed(time.Now().UnixNano())
	population := generateInitialPopulation(6)
	bestFitness := 0
	bestIndividual := Individual{}

	for i := 0; i < 500; i++ {
		population = nextGeneration(population)
		sort.Slice(population, func(i, j int) bool {
			return population[i].fitness() > population[j].fitness()
		})

		if population[0].fitness() > bestFitness {
			bestFitness = population[0].fitness()
			bestIndividual = population[0]
		}
	}

	return bestIndividual
}

func main() {
	var solution Individual
	totalWeight := 0
	totalValue := 0

	for {
		solution = solveKnapsack()
		totalWeight = 0
		totalValue = 0

		for i, bit := range solution.bits {
			if bit == 1 {
				totalWeight += items[i].weight
				totalValue += items[i].value
			}
		}

		if totalWeight <= 400 && totalWeight > 380 {
			break
		}
	}

	fmt.Println("Selected items:")
	for i, bit := range solution.bits {
		if bit == 1 {
			fmt.Printf("%s ", items[i].name)
		}
	}

	fmt.Printf("\nTotal value: %d\nTotal weight: %d\n", totalValue, totalWeight)
}
