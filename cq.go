package main

type Demography struct {
	male    		int
	female			int
	draftable		int
	laborable       int
}

type Food struct {
    pigs			int
	cows			int
	deer            int
	buffaloes       int
	sheep           int
	pigeons			int
	chickens		int
	ducks			int
	geese           int
	snakes			int
}

type Resources struct {
	wood			int
	stone           int
	bronze          int
	copper          int
	iron            int
	straw           int
	cotton          int
	hide            int
	mud             int
	oil             int
}

type Town struct {
	population		Demography
	latitude		int
	longitude       int
	mayor           int
	wealth          int
	food            Food
	supply          Resources
}

KINGDOMS = ["chu", "han", "qi", "qin", "wei", "yan", "zhao"]{}
