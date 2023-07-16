package main

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

func (l *Light) Switch() bool {
	l.on = !l.on
	return l.on
}

func (l *Light) IsOn() bool {
	return l.on
}
