package publisher

// Publisher interface
type Publisher interface {
	GetName() string
	PublishNew()
}
