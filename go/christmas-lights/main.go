package main

import "fmt"

func main() {
	grid := NewGrid(Coordinate{999, 999})

	grid.TurnOn(Coordinate{887, 9}, Coordinate{959, 629})
	grid.TurnOn(Coordinate{454, 398}, Coordinate{844, 448})
	grid.TurnOff(Coordinate{539, 243}, Coordinate{559, 965})
	grid.TurnOff(Coordinate{370, 819}, Coordinate{676, 868})
	grid.TurnOff(Coordinate{145, 40}, Coordinate{370, 997})
	grid.TurnOff(Coordinate{301, 3}, Coordinate{808, 453})
	grid.TurnOn(Coordinate{351, 678}, Coordinate{951, 908})
	grid.Toggle(Coordinate{720, 196}, Coordinate{897, 994})
	grid.Toggle(Coordinate{831, 394}, Coordinate{904, 860})

	fmt.Printf("%d lights remain on\n", grid.LightsOn())
	fmt.Printf("%d brightness\n", grid.Brightness())
}
