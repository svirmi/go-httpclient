package gohttp

// all started with Uppercase letter - exported and publically available outside this package

type whatever struct{}

type HttpClient interface {
	Get() string
	Post()
	Put()
	Patch()
	Delete()
}

// any struct implemented these methods will be concidered HttpClient

// implementations of interface
func (c *whatever) Get() string {}

func (c *whatever) Post() {}

func (c *whatever) Put() {}

func (c *whatever) Patch() {}

func (c *whatever) Delete() {}
