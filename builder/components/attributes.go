package components

// Attributes represents all html attributes.
type Attributes struct {
	Class string            `field:"class"`
	Data  map[string]string `field:"spread"`
}
