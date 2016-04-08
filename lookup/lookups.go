package lookup

//go:generate $GOPATH/bin/stringer -type=Verb,Adjective,Noun

type Verb int
type Adjective int
type Noun int

const (
	arrange Verb = iota
	baste
	beat
	blend
	brown
	bury
	carve
	check
	chop
	close
	cover
	crumple
	cut
	decorate
	discard
	divide
	drape
	drop
	film
	fold
	follow
	form
	force
	glaze
	insert
	lay
	leave
	lift
	make
	melt
	mince
	mix
	moisten
	mound
	open
	pack
	paint
	pierce
	pour
	prepare
	press
	prick
	pull
	puree
	push
	quarter
	raise
	reduce
	refresh
	reheat
	replace
	ring
	roast
	roll
	salt
	saute
	scatter
	scoop
	scrape
	scrub
	season
	separate
	set
	settle
	shave
	simmer
	skim
	slice
	slide
	slip
	slit
	smear
	soak
	spoon
	spread
	sprinkle
	stir
	strain
	strew
	stuff
	surround
	taste
	tie
	tilt
	tip
	top
	toss
	trim
	turn
	twist
	wilt
	wind
	wrap
	LastVerb
)

const (
	acerbic Adjective = iota
	acidic
	acrid
	aged
	ambrosial
	ample
	appealing
	appetizing
	aromatic
	astringent
	baked
	balsamic
	beautiful
	bite_size
	bitter
	bland
	blended
	boiled
	brackish
	briny
	browned
	burnt
	buttered
	caked
	candied
	caramelized
	caustic
	center_cut
	char_broiled
	cheesy
	chilled
	choice
	cholesterol_free
	chunked
	classic
	classy
	clove_coated
	cold
	cool
	copious
	country
	crafted
	creamed
	creamy
	crisp
	crunchy
	cured
	dazzling
	deep_fried
	delectable
	delicious
	delightful
	distinctive
	divine
	doughy
	dressed
	dripping
	drizzled
	dry
	dulcified
	dull
	edible
	elastic
	encrusted
	ethnic
	extraordinary
	famous
	fantastic
	fetid
	fiery
	filet
	fizzy
	flaky
	flat
	flavorful
	flavorless
	flavorsome
	fleshy
	fluffy
	fragile
	free
	free_range
	fresh
	fried
	frosty
	frozen
	fruity
	full
	full_bodied
	furry
	garlicky
	generous
	gingery
	glazed
	golden
	gorgeous
	gourmet
	greasy
	grilled
	gritty
	gustatory
	half
	harsh
	heady
	heaping
	heart_healthy
	heart_smart
	hearty
	heavenly
	homemade
	honeyed
	honey_glazed
	hot
	ice_cold
	icy
	incisive
	indulgent
	infused
	insipid
	intense
	intriguing
	juicy
	jumbo
	kosher
	large
	lavish
	layered
	lean
	leathery
	lemon
	light
	lite
	lightly_salted
	lightly_breaded
	lip_smacking
	lively
	low
	low_sodium
	low_fat
	lukewarm
	luscious
	lush
	marinated
	mashed
	mellow
	mild
	minty
	mixed
	moist
	mouth_watering
	internationally_famous
	nationally_famous
	natural
	nectarous
	non_fat
	nutty
	oily
	organic
	overpowering
	peppery
	petite
	pickled
	piquant
	plain
	pleasant
	plump
	poached
	popular
	pounded
	prepared
	pulpy
	pungent
	pureed
	rancid
	rank
	reduced
	rich
	ripe
	roasted
	robust
	rotten
	rubbery
	saccharine
	saline
	salty
	sapid
	saporific
	saporous
	satin
	satiny
	sauteed
	savorless
	savory
	scrumptious
	sea_salt
	seared
	seasoned
	sharp_tasting
	silky
	simmered
	sizzling
	skillfully
	small
	smelly
	smoked
	smoky
	smooth
	smothered
	soothing
	sour
	southern_style
	special
	spiced
	spicy
	spiral_cut
	spongy
	sprinkled
	stale
	steamed
	steamy
	sticky
	stinging
	strawberry_flavored
	strong
	stuffed
	succulent
	sugar_coated
	sugar_free
	sugared
	sugarless
	sugary
	superb
	sweet
	sweet_and_sour
	sweetened
	syrupy
	tangy
	tantalizing
	tart
	tasteful
	tasteless
	tasty
	tepid
	terrific
	thick
	thin
	toasted
	toothsome
	topped
	tossed
	tough
	traditional
	treacly
	treat
	unflavored
	unsavory
	unseasoned
	vanilla_flavored
	velvety
	vinegary
	warm
	waxy
	weak
	whipped
	whole
	wonderful
	yucky
	yummy
	zesty
	zingy
	boring
	dreamy
	drunk
	goofy
	naughty
	prickly
	sharp
	suspicious
	tender
	LastAdjective
)

//TODO: This needs the RTM scan and removal here and in io.redsift.common.FoodReferenceGenerator
const (
	asparagus Noun = iota
	apple
	apricot
	anchovies
	applesauce
	alphabits_cereal
	artichokes
	animal_crackers
	almonds
	angel_hair_pasta
	acini_de_pepe
	acai_berry
	apple_jacks_cereal
	avocado
	aero_chocolate_bar
	agar_agar
	aubergine
	angel_food_cake
	aloo_gobi
	asiago_cheese
	beef
	beans
	broccoli
	beets
	barley
	bacon
	biscuits
	bundt_cake
	bread
	bagel
	butter
	burrito
	banana
	bologna
	blackberry
	blueberry
	boysenberry
	biscuit
	bun
	barbecue
	bbq
	brownie
	buffalo_wings
	basil
	bok_choy
	bakers_chocolate
	breasts_of_chicken
	brussel_sprouts
	blueberry_muffins
	banana_bread
	bean_salad
	baklava
	buffalo_meat
	black_bean_buns
	balsamic_dressing
	breadfruit
	birds_nest_cookies
	birds_nest_soup
	brains
	bean_sprouts
	black_eyed_peas
	bruschetta
	black_forest_cake
	borscht
	basmati_rice
	bulgur
	butter_squash
	butter_chicken
	corn
	carrots
	cabbage
	cauliflower
	crab
	cracker
	cookies
	cucumber
	croutons
	cantaloupe
	clams
	cranberry_sauce
	corndog
	cake
	cashews
	coconut
	coffee
	casserole
	crepe
	cod
	cumquat
	caramel
	candy
	curry
	cocoa
	cream_of_mushroom
	croissants
	caflib
	chamomile_tea
	corn_chips
	cream_cheese
	cranberries
	currants
	crawfish
	crabcakes
	crème_brule
	cabbage_rolls
	california_dressing
	cassava
	creamsicle
	coroneto
	candy_cane
	coffee_crisp
	corn_flakes
	cream_puff
	cows_tongue
	caviar
	corn_fritters
	cornbread
	couscous
	cobbler
	carrot_cake
	custard
	compote
	coleslaw
	cannoli
	cereal
	cider
	celery
	cinnamon
	caesar_salad
	cilantro
	chips
	cheese
	cherry
	chicken
	chinese_food
	chili_con_carne
	cheeseburger
	cheetos
	cheerios
	chunky_peanut_butter
	chocolate
	chowder
	chow_mien_noodles
	cheddar_cheese
	cheesesteak
	chick_peas
	chocolate_chips
	cheesecake
	chop_suey
	chapatis
	chutney
	churros
	chalupas
	duck
	donut
	danish
	drumsticks
	dutch_apple_pie
	durian
	dairy
	deviled_eggs
	dill
	dill_pickles
	dutch_cookies
	dark_chocolate
	dates
	dim_sum
	dahl_soup
	dream_whip
	dolmades
	dandelion_greens
	ding_dongs
	doritos
	eggs
	eggplant
	endive
	eggroll
	eskimo_pie
	escarole
	eclairs
	escargot
	espresso
	english_muffin
	echinacea
	elbow_macaroni
	eggnog
	enchiladas
	elephant_ears
	eggs_benedict
	edamame
	fish
	french_fries
	fruit
	funnel_cake
	fajitas
	fortune_cookie
	french_toast
	figs
	fudge
	flapjacks
	frijoles
	fritters
	frosting
	fruitcake
	fruit_pie
	frankfurters
	fried_pie
	fig_newtons
	fritos
	fettuccine
	falafal
	free_range_eggs
	feta
	fiddlestick_greens
	flour
	fricassee
	flap_jacks
	flax_seeds
	flax_oil
	focaccia_bread
	fudgicle
	fasolada_soup
	fried_rice
	fruit_salad
	gelatin
	giblets
	ginger
	ginger_ale
	ginger_beer
	gingerbread
	gelato
	garlic
	gravy
	green_beans
	grapes
	gooseberry
	garbonzo
	granola
	grapefruit
	guacamole
	grapeseed_oil
	graham_crackers
	green_olives
	gailan
	green_pepper
	garlic_bread
	goats_milk
	gyoza
	gala_apples
	guava
	gumbo
	gazpacho
	gizzards
	goulash
	ham
	horseradish
	hamburger
	hot_dog
	hoagie
	honeydew
	halibut
	huckleberries
	honey
	hot_chocolate
	heinz_ketchup
	hollandaise
	hummus
	herbs
	herrings
	hearts_of_palm
	hominy
	hash_browns
	haggis
	hog_jowl
	honey_bun
	ho_hos
	ice_cream
	iced_tea
	icee
	icing
	indian_jujubee
	iceberg_lettuce
	italian_bread
	icing_sugar
	indian_curry
	irish_stew
	imitation_baco_bits
	imitation_crab
	impastata_cheese
	indian_pudding
	infant_formula
	instant_coffee
	instant_oatmeal
	italian_dressing
	juice
	jell_o
	jelly
	jam
	jerky
	jalapeño
	jujubes
	jambalaya
	jamaican_patties
	kale
	kohlrabi
	kiwi
	kidney_beans
	kumquat
	kielbasa
	ketchup
	kamut
	lima_bean
	leeks
	lentils
	liver
	lettuce
	lasagna
	lemons
	limes
	lobster
	linguine
	lemonade
	lamb
	lemon_tarts
	long_john_donuts
	legumes
	liverwurst
	licorice
	mustard
	macaroni
	mushroom
	milk
	melon
	muffin
	marshmallows
	mozzarella
	mango
	mince_meat
	marmalade
	mahi_mahi
	maraschino_cherries
	minestrone
	margarine
	maple_syrup
	mayonnaise
	mashed_potatoes
	mackerel
	miso
	manicotti
	martimonial_cake
	mousse
	mars_bar
	meatloaf
	masala
	millet
	moussaka
	mayhaws
	marmite
	noodles
	nuts
	nachos
	nectarines
	nuggets
	nerds_candy
	neapolitan_ice_cream
	nalleys_chips
	naan_bread
	necco_wafers
	neufatual_cheese
	onion
	okra
	orange
	oatmeal
	omelet
	olives
	oysters
	octopus
	oregano
	organic_salad
	ox_tail
	ostrich_eggs
	oat_cakes
	orzo
	pasta
	pickles
	peas
	potato
	potato_chips
	parsley
	pumpkin
	peppers
	parsnips
	pork
	popcorn
	pistachios
	pie
	peanuts
	pizza
	peanut_butter
	pudding
	peaches
	pears
	plum
	prunes
	pancake
	pastry
	pineapple
	peppermint
	pumpernickel
	papaya
	pretzels
	popsicles
	pecans
	pepperoni
	persimmons
	portabello_mushrooms
	pine_nuts
	perogies
	pimentos
	pinto_beans
	paprika
	putanesca
	parmesan
	prawns
	poi
	potstickers
	poutine
	puffed_wheat
	protein_shakes
	protein_bars
	pakora
	passionfruit
	porridge
	plantain
	pilaf
	perch
	pollock
	potroast
	plum_sauce
	pesto
	peking_duck
	paella
	pita_bread
	polenta
	pate
	pavlova
	pancetta
	pomegranate
	quiche
	quail
	quince
	quesadilla
	quinoa
	quaker_oatmeal
	quick_oats
	quakes_rice_cakes
	quorn
	rutabaga
	rhubarb
	radish
	rice
	ribs
	rolls
	rye_bread
	raspberry
	ravioli
	rigatoni
	raisins
	ratatouille
	relish
	ritz_crackers
	rosehips
	romaine
	rooibos
	rice_crackers
	roti
	red_delicious
	rye_crisp
	rice_paper
	rootbeer_float
	red_lentils
	refried_beans
	rissotto
	reuben_sandwich
	ramen
	radicchio
	ricotta_cheese
	romano_cheese
	roll_ups
	spinach
	squash
	soybean
	steak
	spaghetti
	salad
	swiss_cheese
	soup
	sandwich
	sausage
	spam
	scallop
	sushi
	sauce
	succotash
	sundae
	stuffing
	strawberry
	stromboli
	snow_peas
	salami
	stew
	salmon
	sunflower_seeds
	swordfish
	satsuma
	slushies
	squid
	sesame_seeds
	swiss_chard
	sour_cream
	sardines
	salsa
	satziki
	souvlaki
	sashimi
	spelt
	sesame_oil
	spanokopita
	spinach_dip
	sasparilla
	spumoni
	spartan_apples
	starfruit
	sweet_potato
	skors_bars
	smores
	stoned_wheat_crackers
	soysauce
	smoothie_shake
	samosa
	spanish_rice
	snapper
	split_peas
	scones
	stirfry
	sorbet
	sloppy_joes
	string_cheese
	souffle
	starburst
	shish_kebab
	shells
	shepherds_pie
	shrimp
	sherbet
	shallot
	shredded_wheat
	shark_fin_soup
	shortbread
	shortcake
	tomato
	turnip
	turkey
	turtle
	tapioca
	toast
	taco
	tortillas
	taffy
	tarts
	tuna
	tangerine
	tossed_salad
	tabasco_pepper_sauce
	tamale
	tv_dinners
	tater_tots
	tempura
	tofu
	tempeh
	tortellini
	triscuits
	tamarind
	tabbouleh_salad
	teriyaki
	torte
	trifle
	truffle
	turmeric
	tahini
	twinkies
	twizzlers
	tostitos
	thighs
	thousand_island_salad_dressing
	thimbleberry
	upside_down_cake
	uglifruit
	unpasteurized_milk
	vegetables
	venison
	veal
	vermicelli
	vidalia_onion
	vinegar
	vol_au_vent
	vegemite
	vichyssoise
	vanilla
	vine_leaves
	vegetable_soup
	virgin_oil
	vakalolo
	vindaloo
	vomit_fruit
	watercress
	wax_bean
	watermelon
	wheat_bread
	wings
	waffles
	walnuts
	wafers
	walleye
	water
	wontons
	wonder_bread
	wasabi
	wild_rice
	waldorf_salad
	xmas_cake
	xigua
	chex_mix
	yam
	yellow_squash
	yogurt
	yolk
	yellow_beans
	yorkshire_pudding
	yucca
	yautia
	yaki_soba
	yiyantes
	zucchini
	ziti
	zwieback_bread
	zabaglione_sauce
	zander_fish
	zours
	LastNoun
)