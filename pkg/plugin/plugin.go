package plugin

// YomoObjectPlugin defines the interface of a normal plugin which
// handles object data
type YomoObjectPlugin interface {
	Handle(value interface{}) (interface{}, error)
	Observed() string
	Mold() interface{}
	Name() string
}

// YomoStreamPlugin defines the interface of a stream-oritened plugin
type YomoStreamPlugin interface {
	HandleStream(buf []byte, done bool) ([]byte, error)
	Observed() string
	Name() string
}
