package service

import (
	"net/http"

	"github.com/rituK/com/ritu/controller"
)

// Defines a single route, e.g. a human readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{

		"GetUsers",      // Name
		"GET",           // HTTP method
		"/api/getUsers", // Route pattern
		controller.GetUsers,
	},
	Route{

		"InsertUsers",     // Name
		"POST",            // HTTP method
		"/api/createUser", // Route pattern
		controller.InsertUser,
	},
	Route{

		"GetUser",           // Name
		"GET",               // HTTP method
		"/api/getUser/{id}", // Route pattern
		controller.GetUser,
	},
	Route{

		"UpdateUser",           // Name
		"POST",                 // HTTP method
		"/api/updateUser/{id}", // Route pattern
		controller.UpdateUser,
	},
	Route{

		"DeleteUser",           // Name
		"Delete",               // HTTP method
		"/api/removeUser/{id}", // Route pattern
		controller.DeleteUser,
	},
}
