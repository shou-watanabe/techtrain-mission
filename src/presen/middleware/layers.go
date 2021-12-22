package middleware

import (
	"net/http"
)

func Layers(handler http.Handler) http.Handler {
	return Recovery(
		Logger(
			handler,
		),
	)
}

// func AuthLayers(handler http.Handler) http.Handler {
// 	return Recovery(
// 		Context(
// 			Access(
// 				Basic(
// 					handler,
// 				),
// 			),
// 		),
// 	)
// }
