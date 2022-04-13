package config

func (app *Server) RegisterRouter() {
	// Register routes
	app.engine.Get("/", app.handlers.Index)
	app.engine.Post("/topsecret", app.handlers.TopSecret)
	app.engine.Post("/topsecret_split/:satilliteName", app.handlers.TopScretSplitPost)
	app.engine.Get("/topsecret_split/:satilliteName", app.handlers.TopScretSplitGet)
}
