package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)




func (app *application) routes()  *httprouter.Router {
	router := httprouter.New()


	router.HandlerFunc(http.MethodGet,"/",app.Home);

	return router
}