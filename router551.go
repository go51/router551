package router551

import (
	"github.com/go51/container551"
	"github.com/go51/string551"
	"regexp"
	"runtime"
)

type routerMethod int

const (
	GET     routerMethod = 1
	POST    routerMethod = 2
	PUT     routerMethod = 4
	DELETE  routerMethod = 8
	COMMAND routerMethod = 16
	UNKNOWN routerMethod = 32
)

func (rm routerMethod) String() string {
	switch rm {
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

type ActionFunc func(c *container551.Container) interface{}

type Router struct {
	get     map[string]*route
	post    map[string]*route
	put     map[string]*route
	delete  map[string]*route
	command map[string]*route
}

var routerInstance *Router

func Load() *Router {
	if routerInstance != nil {
		return routerInstance
	}

	routerInstance = &Router{
		get:     make(map[string]*route, 0),
		post:    make(map[string]*route, 0),
		put:     make(map[string]*route, 0),
		delete:  make(map[string]*route, 0),
		command: make(map[string]*route, 0),
	}

	return routerInstance
}

type route struct {
	name        string
	pattern     string
	regex       *regexp.Regexp
	packageName string
	action      ActionFunc
	keys        []string
}

func (r *route) Name() string {
	return r.name
}

func (r *route) Keys() []string {
	return r.keys
}

func (r *route) Action() ActionFunc {
	return r.action
}

func (r *route) PackageName() string {
	return r.packageName
}

func (r *Router) Add(method routerMethod, name, pattern string, action ActionFunc) {

	pc, _, _, _ := runtime.Caller(1)
	path := runtime.FuncForPC(pc).Name()

	route := &route{
		name:        name,
		pattern:     pattern,
		regex:       nil,
		packageName: r.getPackageName(path),
		action:      action,
		keys:        []string{},
	}

	if method != COMMAND {
		route.keys = r.getKeys(pattern)
		regexPattern := pattern
		for i := 0; i < len(route.keys); i++ {
			regexPattern = string551.Replace(regexPattern, ":"+route.keys[i]+":", ".*")
		}
		route.regex = regexp.MustCompile(regexPattern)
	}

	if method&GET == GET {
		r.addGet(name, route)
	}
	if method&POST == POST {
		r.addPost(name, route)
	}
	if method&PUT == PUT {
		r.addPut(name, route)
	}
	if method&DELETE == DELETE {
		r.addDelete(name, route)
	}
	if method&COMMAND == COMMAND {
		r.addCommand(name, route)
	}
}

func (r *Router) addGet(name string, route *route) {
	r.get[name] = route
}

func (r *Router) addPost(name string, route *route) {
	r.post[name] = route
}

func (r *Router) addPut(name string, route *route) {
	r.put[name] = route
}

func (r *Router) addDelete(name string, route *route) {
	r.delete[name] = route
}

func (r *Router) addCommand(name string, route *route) {
	r.command[name] = route
}

func (r *Router) getKeys(pattern string) []string {
	keys := []string{}

	patternBytes := string551.StringToBytes(pattern)
	coron := false

	keyBytes := []byte{}

	for i := 0; i < len(patternBytes); i++ {
		if patternBytes[i] == 0x3A { // 0x3A => ":"
			if !coron {
				keyBytes = []byte{}
				coron = true
			} else {
				keys = append(keys, string551.BytesToString(keyBytes))
				coron = false
			}
		} else {
			if coron {
				keyBytes = append(keyBytes, patternBytes[i])
			}
		}
	}

	return keys
}

var regexpControllerPattern *regexp.Regexp

func (r *Router) getPackageName(path string) string {
	if regexpControllerPattern == nil {
		regexpControllerPattern = regexp.MustCompile(`\/\w+\.`)
	}

	packageName := regexpControllerPattern.FindString(path)

	return packageName[1 : len(packageName)-1]
}

func (r *Router) FindRouteByName(method, name string) *route {
	if method == GET.String() {
		return r.get[name]
	}
	if method == POST.String() {
		return r.post[name]
	}
	if method == PUT.String() {
		return r.put[name]
	}
	if method == DELETE.String() {
		return r.delete[name]
	}
	if method == COMMAND.String() {
		return r.command[name]
	}

	return nil
}

func (r *Router) FindRouteByPathMatch(method, path string) *route {
	routes := r.getRoutes(method)

	for _, route := range routes {
		if route.pattern == path {
			return route
		}
	}

	for _, route := range routes {
		if route.pattern == "/" {
			continue
		}
		if route.regex.MatchString(path) {
			return route
		}
	}

	return nil
}

func (r *Router) getRoutes(method string) map[string]*route {
	switch method {
	case GET.String():
		return r.get
	case POST.String():
		return r.post
	case PUT.String():
		return r.put
	case DELETE.String():
		return r.delete
	case COMMAND.String():
		return r.command
	}

	return make(map[string]*route)
}
