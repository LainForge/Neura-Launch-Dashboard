package server

func(server *Server) PingRoute(){
	server.Router.GET("/", server.Ping)
}