package router551

type method int

const (
	GET method = iota
	POST
	PUT
	DELETE
	COMMAND
	UNKNOWN
)

func (m method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case COMMAND:
		return "COMMAND"
	}
	return "UNKNOWN"
}

type Router struct{}

var routerInstance *Router

func Load() *Router {
	if routerInstance != nil {
		return routerInstance
	}

	routerInstance = &Router{}

	return routerInstance
}
