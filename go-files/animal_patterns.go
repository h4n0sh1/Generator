package main

import (
	"image/color"
	"math/rand/v2"
)

// AnimalPattern represents a pixel art pattern for an animal
type AnimalPattern struct {
	Name   string
	Width  int
	Height int
	// Pattern uses characters: ' ' = transparent, '#' = primary color, 'o' = secondary color, '.' = accent
	Pattern   []string
	Primary   color.RGBA
	Secondary color.RGBA
	Accent    color.RGBA
}

// GetRandomAnimal returns a random animal pattern with randomized colors
func GetRandomAnimal() AnimalPattern {
	animals := []AnimalPattern{
		catPattern(),
		dogPattern(),
		birdPattern(),
		fishPattern(),
		bunnyPattern(),
		bearPattern(),
		foxPattern(),
		owlPattern(),
		penguinPattern(),
		elephantPattern(),
	}

	animal := animals[rand.IntN(len(animals))]

	// Randomize colors for variety
	animal.Primary = randomBrightColor()
	animal.Secondary = randomBrightColor()
	animal.Accent = color.RGBA{0, 0, 0, 255} // Keep accent as black for eyes/details

	return animal
}

func randomBrightColor() color.RGBA {
	colors := []color.RGBA{
		{255, 107, 107, 255}, // Red
		{255, 159, 67, 255},  // Orange
		{254, 202, 87, 255},  // Yellow
		{29, 209, 161, 255},  // Green
		{72, 219, 251, 255},  // Cyan
		{84, 160, 255, 255},  // Blue
		{156, 136, 255, 255}, // Purple
		{255, 107, 181, 255}, // Pink
		{162, 155, 254, 255}, // Lavender
		{46, 213, 115, 255},  // Lime
		{255, 165, 2, 255},   // Gold
		{255, 127, 80, 255},  // Coral
	}
	return colors[rand.IntN(len(colors))]
}

func catPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Cat",
		Width:  16,
		Height: 16,
		Pattern: []string{
			"  ##      ##  ",
			" ####    #### ",
			" ############",
			"##############",
			"##.##oooo##.##",
			"##############",
			"######oo######",
			" ####oooo#### ",
			"  ##########  ",
			"    ######    ",
			"   ########   ",
			"  ##########  ",
			" ############ ",
			" ##  ####  ## ",
			" ##  ####  ## ",
			"     ####     ",
		},
	}
}

func dogPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Dog",
		Width:  16,
		Height: 16,
		Pattern: []string{
			"  ####  ####  ",
			" ###### ######",
			" ###### ######",
			"  ########### ",
			" ############ ",
			"##.##oooo##.##",
			"##############",
			"######oo######",
			" ####oooo#### ",
			"  ##oooooo##  ",
			"  ##########  ",
			"  ##########  ",
			" ############ ",
			" ##  ####  ## ",
			" ##  ####  ## ",
			"     ####     ",
		},
	}
}

func birdPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Bird",
		Width:  14,
		Height: 12,
		Pattern: []string{
			"     ####     ",
			"    ######    ",
			"   ##.#####   ",
			"   ########   ",
			"  ooo#######  ",
			"   #########  ",
			"    ########  ",
			"     ######   ",
			"      ####    ",
			"     ######   ",
			"    ##    ##  ",
			"    #      #  ",
		},
	}
}

func fishPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Fish",
		Width:  16,
		Height: 10,
		Pattern: []string{
			"       ####    ",
			"  ## ########  ",
			" ####.#########",
			"################",
			"################",
			"################",
			" ####.#########",
			"  ## ########  ",
			"       ####    ",
			"               ",
		},
	}
}

func bunnyPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Bunny",
		Width:  14,
		Height: 16,
		Pattern: []string{
			"   ##    ##   ",
			"  ####  ####  ",
			"  ####  ####  ",
			"  ####  ####  ",
			" ####oo####   ",
			" ############ ",
			"##.##oooo##.##",
			"##############",
			"######oo######",
			" ############ ",
			"  ##########  ",
			"  ##########  ",
			" ############ ",
			"    ##  ##    ",
			"    ##  ##    ",
			"              ",
		},
	}
}

func bearPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Bear",
		Width:  16,
		Height: 16,
		Pattern: []string{
			" ###      ### ",
			"#####    #####",
			"#####    #####",
			" ##########   ",
			"##############",
			"##############",
			"##.##oooo##.##",
			"##############",
			"######oo######",
			" ####oooo#### ",
			"  ##########  ",
			"  ##########  ",
			" ############ ",
			" ##  ####  ## ",
			"####  ##  ####",
			"              ",
		},
	}
}

func foxPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Fox",
		Width:  16,
		Height: 16,
		Pattern: []string{
			"##          ##",
			"###        ###",
			"####      ####",
			"#####oooo#####",
			"##oooooooooo##",
			"#.#oooooooo#.#",
			"##oooooooooo##",
			" ##ooo..ooo## ",
			"  ##oooooo##  ",
			"   ########   ",
			"   ########   ",
			"  ##########  ",
			" ############ ",
			" ##  ####  ## ",
			" ##  ####  ## ",
			"              ",
		},
	}
}

func owlPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Owl",
		Width:  14,
		Height: 16,
		Pattern: []string{
			"  ##      ##  ",
			" ####    #### ",
			"##############",
			"##############",
			"###oo####oo###",
			"##.oo####oo.##",
			"###oo####oo###",
			"######oo######",
			" ####oooo#### ",
			" ####oooo#### ",
			"  ##########  ",
			"  ##########  ",
			"  ##########  ",
			"  ###    ###  ",
			"  ##      ##  ",
			"              ",
		},
	}
}

func penguinPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Penguin",
		Width:  12,
		Height: 16,
		Pattern: []string{
			"    ####    ",
			"   ######   ",
			"  ########  ",
			"  ##oooo##  ",
			" ##.oo.oo## ",
			" ###oooo### ",
			" ##oooooo## ",
			"  #oooooo#  ",
			"  #oooooo#  ",
			"  #oooooo#  ",
			"  ##oooo##  ",
			"  ##oooo##  ",
			"   ##oo##   ",
			"  ###  ###  ",
			" ###    ### ",
			"            ",
		},
	}
}

func elephantPattern() AnimalPattern {
	return AnimalPattern{
		Name:   "Elephant",
		Width:  16,
		Height: 14,
		Pattern: []string{
			"   ########   ",
			"  ##########  ",
			" ############ ",
			" ##.######.## ",
			"##############",
			"##############",
			"###  ####  ###",
			"###  ####  ###",
			"     ####     ",
			"     ####     ",
			"    ######    ",
			"   ##    ##   ",
			"  ###    ###  ",
			" ####    #### ",
		},
	}
}
