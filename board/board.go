package Board

const (
	GROUP_PURPLE = iota
	GROUP_LITE_BLUE
	GROUP_PINK
	GROUP_ORANGE
	GROUP_RED
	GROUP_YELLOW
	GROUP_GREEN
	GROUP_BLUE
)

type Property struct {
	BoardPiece

	name string
	group int32
	rates [6]int32
	mortgage int32
	house_cost int32

	//owner *Player
	is_mortaged bool
}


type BoardPiece interface {
	PlayerPass()
	PlayerLand()
}

type BoardGo struct {
	BoardPiece

	reward int32
}

type CommunityChest struct {
	BoardPiece
}

type Chance struct {
	BoardPiece
}

type IncomeTax struct {
	BoardPiece
}

type Jail struct {
	BoardPiece
}

type GotoJail struct {
	BoardPiece
}

type Railroad struct {
	BoardPiece

	name string
}

type FreeParking struct {
	BoardPiece
}

type ElectricCompany struct {
	BoardPiece
}

type WaterWorks struct {
	BoardPiece
}

/* Game board definations based off of pre-2008 U.S. Standard Edition */
var Spaces = []BoardPiece { 
	BoardGo{ reward: 200 },
	Property{
		name: 		"Mediterranean Avenue",
		group: 		GROUP_PURPLE,
		rates: 		[6]int32{ 2, 10, 30, 90, 160, 250 }, 
		mortgage: 	30, 
		house_cost: 50,
	},
	CommunityChest{},
	Property {
		name: 		"Baltic Avenue",
		group: 		GROUP_PURPLE,
		rates: 		[6]int32{ 4, 20, 60, 180, 320, 450 }, 
		mortgage: 	30, 
		house_cost: 50,
	},
	IncomeTax{},
	Railroad { name: "Reading Railroad" },
	Property {
		name: 		"Oriental Avenue",
		group: 		GROUP_LITE_BLUE,
		rates: 		[6]int32{ 6, 30, 90, 270, 400, 550 }, 
		mortgage: 	50, 
		house_cost: 50,
	},
	Chance { },
	Property {
		name: 		"Vermont Avenue",
		group: 		GROUP_LITE_BLUE,
		rates: 		[6]int32{ 6, 30, 90, 270, 400, 550 },
		mortgage: 	50,
		house_cost: 50,
	},
	Property{
		name: 		"Connecticut Avenue", 
		group: 		GROUP_LITE_BLUE,
		rates: 		[6]int32{ 8, 40, 100, 300, 450, 600 }, 
		mortgage: 	60, 
		house_cost: 50,
	},
	Jail{},
	/* Left */
	Property{
		name: 		"St. Charles Place",
		group: 		GROUP_PINK,
		rates: 		[6]int32{ 10, 50, 150, 450, 625, 750 }, 
		mortgage: 	70, 
		house_cost: 100,
	},
	ElectricCompany{},
	Property{
		name: 		"States Avenue",
		group: 		GROUP_PINK,
		rates: 		[6]int32{ 10, 50, 150, 450, 625, 750 }, 
		mortgage: 	70, 
		house_cost: 100,
	},
	Property{
		name: 		"Virginia Avenue",
		group: 		GROUP_PINK,
		rates: 		[6]int32{ 12, 60, 180, 500, 700, 900 }, 
		mortgage: 	80, 
		house_cost: 100,
	},
	Railroad{ name: "Pennsylvania Railroad" },
	Property{
		name:		"St. James Place",
		group: 		GROUP_ORANGE,
		rates: 		[6]int32{ 14, 70, 200, 550, 750, 950 }, 
		mortgage: 	90, 
		house_cost: 100,
	},
	CommunityChest{},
	Property{
		name:		"Tennessee Avenue",
		group: 		GROUP_ORANGE,
		rates: 		[6]int32{ 14, 70, 200, 550, 750, 950 }, 
		mortgage: 	90,
		house_cost: 100,
	},
	Property{
		name: 		"New York Avenue",
		group: 		GROUP_ORANGE,
		rates: 		[6]int32{ 16, 80, 220, 600, 800, 1000 }, 
		mortgage: 	100,
		house_cost: 100,
	},
	/* Top */
	FreeParking{},
	Property{
		name:		"Kentucky Avenue",
		group: 		GROUP_RED,
		rates: 		[6]int32{ 18, 90, 250, 700, 875, 1050 }, 
		mortgage: 	110,  
		house_cost: 150,
	},
	Chance{},
	Property{
		name:		"Indiana Avenue",
		group: 		GROUP_RED,
		rates: 		[6]int32{ 18, 90, 250, 700, 875, 1050, }, 
		mortgage: 	110, 
		house_cost: 150,
	},
	Property{
		name:		"Illinois Avenue",
		group: 		GROUP_RED,
		rates: 		[6]int32{ 20, 100, 300, 750, 925, 1100 }, 
		mortgage: 	120, 
		house_cost: 150,
	},
	Railroad{ name: "B. & O. Railroad" },
	Property{
		name: 		"Atlantic Avenue",
		group: 		GROUP_YELLOW,
		rates: 		[6]int32{ 22, 110, 330, 800, 975, 1150 }, 
		mortgage: 	130, 
		house_cost: 150,
	},
	Property{
		name: 		"Ventnor Avenue",
		group: 		GROUP_YELLOW,
		rates: 		[6]int32{ 22, 110, 330, 800, 975, 1150 }, 
		mortgage: 	130, 
		house_cost: 150,
	},
	WaterWorks{},
	Property{
		name: 		"Marvin Gardens",
		group: 		GROUP_YELLOW,
		rates: 		[6]int32{ 24, 120, 360, 850, 1025, 1200 }, 
		mortgage: 	140, 
		house_cost: 150,
	},
	GotoJail{},
	/* Right */
	Property{
		name: 		"Pacific Avenue",
		group: 		GROUP_GREEN,
		rates: 		[6]int32{ 26, 130, 390, 900, 1100, 1275 }, 
		mortgage: 	150, 
		house_cost: 200,
	},
	Property{
		name: 		"North Carolina Avenue",
		group: 		GROUP_GREEN,
		rates: 		[6]int32{ 26, 130, 390, 900, 1100, 1275 }, 
		mortgage: 	150, 
		house_cost: 200,
	},
	CommunityChest{},
	Property{
		name: 		"Pennsylvania Avenue",
		group: 		GROUP_GREEN,
		rates: 		[6]int32{ 28, 150, 450, 1000, 1200, 1400 }, 
		mortgage: 	160, 
		house_cost: 200,
	},
	Railroad{name: "Short Line"},
	Chance{},
	Property{
		name: 		"Park Place",
		group: 		GROUP_BLUE,
		rates: 		[6]int32{ 35, 175, 500, 1100, 1300, 1500 }, 
		mortgage: 	175, 
		house_cost: 200,
	},
	Property{
		name: 		"Boardwalk",
		group: 		GROUP_BLUE,
		rates: 		[6]int32{ 50, 200, 600, 1400, 1700, 2000 }, 
		mortgage: 	200, 
		house_cost: 200,
	},
}
