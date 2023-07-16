package main

type Grid struct {
	lights      [][]*Light
	nOfLightsOn int
	brightness  int
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

func (g *Grid) Brightness() int {
	return g.brightness
}

func (g *Grid) TurnOn(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			light := g.lights[x][y]
			g.turnOnAndCount(light)
		}
	}
}

func (g *Grid) turnOnAndCount(light *Light) {
	if switched := light.TurnOn(); switched {
		g.nOfLightsOn++
	}

	g.increaseBrightnessBy(1)
}

func (g *Grid) TurnOff(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			light := g.lights[x][y]
			g.turnOffAndCount(light)
		}
	}
}

func (g *Grid) turnOffAndCount(light *Light) {
	if switched := light.TurnOff(); switched {
		g.nOfLightsOn--
	}

	g.increaseBrightnessBy(-1)
}

func (g *Grid) Toggle(from, to Coordinate) {
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			light := g.lights[x][y]
			g.toggleAndCount(light)
		}
	}
}

func (g *Grid) toggleAndCount(light *Light) {
	g.increaseBrightnessBy(2)

	if light.Switch() {
		g.nOfLightsOn++
		return
	}
	g.nOfLightsOn--

}

func (g *Grid) increaseBrightnessBy(amount int) {
	g.brightness += amount
}
