package main

var Players = []Player { 
	Player { 
		name: "Dustin", 
		cash: 2000,
		position: 0,
	},
}

var CurrentPlayer = &Players[0]