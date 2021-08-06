package main

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
}

func main() {
}
