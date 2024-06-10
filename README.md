This program solves the knapsack problem using a genetic algorithm. The knapsack problem involves selecting a combination of items to maximize their value in a knapsack with limited capacity (maximum weight). The genetic algorithm takes an evolutionary approach to finding the optimal solution.

The program implements the following steps:

Initialization of a random population.
Computing the fitness of each individual in the population, where fitness is measured as the total value of the items in the knapsack.
Selection of parents for creating the next generation using tournament selection.
Generating a new generation through chromosome crossover and mutation.
Repeating steps 2-4 for a certain number of generations or until a stop condition is met.
The program returns a set of items that maximize the value in the knapsack without exceeding the maximum weight. If the maximum weight of the knapsack is exceeded, the program returns the best result obtained up to that point.

I'll attach a screenshot of the program in action, where it selects the optimal combination of items for the knapsack and displays the total value and weight of these items.
![image](https://github.com/dmytroserhiienko/backpacks/assets/73895498/eb6cdf20-4bf8-46d4-81a9-5dd399a0a6a4)
