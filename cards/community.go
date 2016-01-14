package Cards

type CommunityCard struct {
	name string
	holdable bool
	action func()
}

var CommunityCards = []CommunityCard { 
	CommunityCard { 
		name: "Income Tax Refund Collect $20",
		action : func() {
			//CurrentPlayer.cash += 20
		},
	},
	CommunityCard { 
		name: "You are assessed for street repairs. $40 per house $115 per hotel",
		action : func() {

		},
	},
	CommunityCard { 
		name: "You inherit &100!",
		action : func() {
			//CurrentPlayer.cash += 10
		},
	},
	CommunityCard { 
		name: "Grand Opera Opening. Collect $50 from every player",
		holdable: false,
		action : func() {

		},
	},
	CommunityCard { 
		name: "Xmas fund matures Collect $100",
		action : func() {
			//CurrentPlayer.cash += 100
		},
	},
	CommunityCard { 
		name: "Advance to Go Collect $200",
		action : func() {

		},
	},
	CommunityCard { 
		name: "Bank Error in your favor Collect $200",
		action : func() {
			//CurrentPlayer.cash += 200
		},
	},
	CommunityCard { 
		name: "Get out of jail free card",
		holdable: true,
		action : func() {

		},
	},
	CommunityCard { 
		name: "Pay hospital $100",
		action : func() {
			//CurrentPlayer.cash -= 100
		},
	},
	CommunityCard { 
		name: "Receive for Services $25",
		action : func() {
			//CurrentPlayer.cash += 25
		},
	},
	CommunityCard { 
		name: "Go to to Jail",
		action : func() {

		},
	},
	CommunityCard { 
		name: "Pay school tax of $150",
		action : func() {
			//CurrentPlayer.cash -= 150
		},
	},
	CommunityCard {
		name: "Doctors Fee Pay $50",
		action : func() {
			//CurrentPlayer.cash -= 50
		},
	},
	CommunityCard { 
		name: "From sale of stock You get $45",
		action : func() {
			//CurrentPlayer.cash += 45
		},
	},
	CommunityCard { 
		name: "Life insurance matures Collect $100",
		action : func() {
			//CurrentPlayer.cash += 100
		},
	},
	CommunityCard { 
		name: "You have won second prize in a beauty contest! Collect $10",
		action : func() {
			//CurrentPlayer.cash += 10
		},
	},
}