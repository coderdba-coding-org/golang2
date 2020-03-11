package agent

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Agent struct {
	//datastore      *datastore.Datastore
	router         *mux.Router
	//waitconditions *memorystore.MemoryStore
	///sniffer        *dhcpsniffer.DhcpSniffer
	//vmIpMap        *inMemoryVmIpMap
}

type VersionedRoutes map[string][]Route

type Route struct {
	RouteName         string
	Methods           string
	Pattern           string
	QueryList         []string
	HandlerFunction   http.HandlerFunc
}

// functions

// Main function that starts the agent
func NewAgent() (*Agent, error) {

  a := &Agent{}  // create an empty Agent

  // build the Agent
  a.router = NewRouter(a.getVersionedRoutes())

}

// Creates a new router
func NewRouter(versionedRoutes VersionedRoutes) *mux.Router {

  // get a new router
	router := mux.NewRouter().StrictSlash(true)

  // add routes to the router
	for version, routes := range versionedRoutes {
		for _, route := range routes {

			generatedPath := generateVersionedPath(version, route.Pattern)

			handler := LogRequest(route.HandlerFunc)
			if route.Authenticated {
				handler = CheckAuthenticated(LogRequest(route.HandlerFunc))
			}
			router.
				Methods(route.Method).
				Path(generatedPath).
				Queries(route.Queries...).
				Name(route.Name).
				Handler(handler)
		}
	}
	return router
}

// Get versioned-routes list
