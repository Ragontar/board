package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = CORSheaders(handler)

		router.
			Methods(route.Method, http.MethodOptions).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"ChannelListUserIdGet",
		strings.ToUpper("Get"),
		"/channel-list/{user-id}",
		ChannelListUserIdGet,
	},

	Route{
		"GroupListUserIdCategoryCategoryIdPut",
		strings.ToUpper("Put"),
		"/group-list/{user-id}/category/{category-id}",
		GroupListUserIdCategoryCategoryIdPut,
	},

	Route{
		"GroupListUserIdCategoryPost",
		strings.ToUpper("Post"),
		"/group-list/{user-id}/category",
		GroupListUserIdCategoryPost,
	},

	Route{
		"GroupListUserIdGet",
		strings.ToUpper("Get"),
		"/group-list/{user-id}",
		GroupListUserIdGet,
	},

	Route{
		"GroupListUserIdGroupGroupIdPut",
		strings.ToUpper("Put"),
		"/group-list/{user-id}/group/{group-id}",
		GroupListUserIdGroupGroupIdPut,
	},

	Route{
		"GroupListUserIdGroupPost",
		strings.ToUpper("Post"),
		"/group-list/{user-id}/group",
		GroupListUserIdGroupPost,
	},

	Route{
		"LinkTelegramUserIdConfirmPut",
		strings.ToUpper("Put"),
		"/link/telegram/{user-id}/confirm",
		LinkTelegramUserIdConfirmPut,
	},

	Route{
		"LinkTelegramUserIdPut",
		strings.ToUpper("Put"),
		"/link/telegram/{user-id}",
		LinkTelegramUserIdPut,
	},

	Route{
		"MessagesGroupGroupIdGet",
		strings.ToUpper("Get"),
		"/messages/group/{group-id}",
		MessagesGroupGroupIdGet,
	},

	Route{
		"AuthorizationPost",
		strings.ToUpper("Post"),
		"/authorization",
		AuthorizationPost,
	},

	Route{
		"RegistrationPost",
		strings.ToUpper("Post"),
		"/registration",
		RegistrationPost,
	},
}
