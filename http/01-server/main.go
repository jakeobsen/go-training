package main

// test this with: curl localhost:8000/test -d '[{"Name": "Platypus", "Order": "Monotremata"}]'

func main() {
	server := newServer("0.0.0.0", "8000")
	server.serve()
}
