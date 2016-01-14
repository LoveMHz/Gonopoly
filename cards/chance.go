package Cards

type ChanceCard struct {
	name string
	holdable bool
	action func()
}

var ChanceCards = []ChanceCard { 
	ChanceCard{ 
		name: "Take a ride on the Reading Railroad. If you pass go collect $200",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Bank pays you dividend of $50",
		action : func() {
			//CurrentPlayer.cash += 50
		},
	},
	ChanceCard{ 
		name: "Advance to Illinois Avenue",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Your building and loan matures. Collect $150",
		action : func() {
			//CurrentPlayer.cash += 150
		},
	},
	ChanceCard{ 
		name: "Get out of jail free card",
		holdable: true,
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Make general repairs on all your property. Pay $25 for each house. Pay $100 " +
			  "for each hotel.",
		action: func() {

		},
	},
	ChanceCard{ 
		name: "Advance token to the nearest railroad and pay the owner twice the rental to " + 
			  "which he is otherwise entitled. If railroad is unowned, you may buy it from " + 
			  "the bank.",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Pay poor tax of $15",
		action : func() {
			//CurrentPlayer.cash -= 50
		},
	},
	ChanceCard{ 
		name: "Take a walk on the Boardwalk",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Advance to St. Charles Place",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "You have been elected chairman of the board. Pay each player $50",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Advance token to nearest utility. If unowned, you may buy it from the bank. " + 
			  "If owned, throw the dice and pay owner a total of 10 times the amount thrown",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Go back 3 spaces",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Advance to Go. Collect $200 dollars",
		action : func() {

		},
	},
	ChanceCard{ 
		name: "Go directly to jail",
		action : func() {

		},
	},
}