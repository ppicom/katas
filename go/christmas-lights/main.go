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
}

type Coordinate struct {
	X int
	Y int
}

type Light struct {
	on bool
}

func (l *Light) TurnOn() (switched bool) {
	if l.on == false {
		switched = true
	}
	l.on = true
	return
}

func (l *Light) TurnOff() (switched bool) {
	if l.on == true {
		switched = true
	}
	l.on = false
	return
}

func (l *Light) IsOn() bool {
	return l.on
}

type Grid struct {
	lights      [][]*Light
	nOfLightsOn int
}

func NewGrid(size Coordinate) Grid {
	lights := make([][]*Light, size.X+1)
	for row := range lights {
		lights[row] = make([]*Light, size.Y+1)
		for l := range lights[row] {
			lights[row][l] = &Light{}
		}
	}

	return Grid{
		lights: lights,
	}
}

func (g *Grid) LightsOn() int {
	return g.nOfLightsOn
}

func (g *Grid) TurnOn(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			if switched := g.lights[x][y].TurnOn(); switched {
				g.nOfLightsOn += 1
			}
		}
	}
}

func (g *Grid) TurnOff(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			if switched := g.lights[x][y].TurnOff(); switched {
				g.nOfLightsOn -= 1
			}
		}
	}
}

func (g *Grid) Toggle(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			l := g.lights[x][y]
			if l.IsOn() {
				l.TurnOff()
				g.nOfLightsOn -= 1
			} else {
				l.TurnOn()
				g.nOfLightsOn += 1
			}
		}
	}
}
