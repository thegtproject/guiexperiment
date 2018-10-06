package ui

type EventDispatcher interface {
	Subscribe(Event)
}

type MouseEvent struct{}
