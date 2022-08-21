package snake

// CommandID represents direction
type CommandID int

const (
	UP CommandID = iota
	DOWN
	LEFT
	RIGHT

	ESC
	QUIT
)

// Controller interface for controller of snake
type Controller interface {
	Control(chan<- CommandID)
}
