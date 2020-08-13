package main

func main() {
	server := newServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/home", HandleHome)
	server.Handle("POST", "/api", server.addMiddleware(HandleHome, checkAuthentication(), logger()))
	server.Handle("POST", "/create", postRequest)
	server.Handle("POST", "/user", userPostRequest)

	server.listen()
}
