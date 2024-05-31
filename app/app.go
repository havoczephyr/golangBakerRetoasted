package app

import (
	"fmt"
	"math"
	"strings"
)

func App(values ImportValues) string {
	recipeName := "Shokupan Bread"
	cold := 0.29
	coldDensity := 1.0
	warm := 0.25
	warmDensity := 0.99
	var flourPercents []flourPercent
	var ingredientPercents []ingredientPercent
	var floursMass []flourMass
	var ingredientsMass []ingredientMass
	softwareAuthor := "mhz"
	description := "Shokupan is a buttery japanese style milkbread with a flavorful consistency. Shokupan has a dense and flavorful but very soft inner crumb and a crunchy buttery exterior."
	var documentBuilder strings.Builder

	volumeRatio := cold
	density := coldDensity

	if values.TempAtmosphere {
		volumeRatio = warm
		density = warmDensity
	}

	flourPercents = append(flourPercents,
		flourPercent{"Bread Flour", 1.0})

	ingredientPercents = append(ingredientPercents,
		ingredientPercent{"Sugar", 0.048},
		ingredientPercent{"Salt", 0.020},
		ingredientPercent{"Yeast", 0.012},
		ingredientPercent{"Milk", 0.718},
		ingredientPercent{"Butter", 0.054})

	containerVolume := values.Length * values.Width * values.Depth
	doughVolume := containerVolume * volumeRatio
	doughMass := doughVolume * density
	var totalPercentage float64

	for _, flour := range flourPercents {
		totalPercentage += flour.percentageValue
	}
	for _, ingredient := range ingredientPercents {
		totalPercentage += ingredient.percentageValue
	}

	flour_total := doughMass / totalPercentage

	for _, ingredient := range ingredientPercents {
		ingredientsMass = append(ingredientsMass, ingredientMass{ingredient.ingredientName, flour_total * ingredient.percentageValue})
	}
	for _, flour := range flourPercents {
		floursMass = append(floursMass, flourMass{flour.flourName, flour_total * flour.percentageValue})

	}

	documentBuilder.WriteString(fmt.Sprintf("# %s\n### Recipe by %s \n\n%s\n\n**Flours**:", recipeName, softwareAuthor, description))

	for _, flour := range floursMass {
		documentBuilder.WriteString(fmt.Sprintf("\n- %s: %.0fg", flour.flourName, math.Round(flour.massValue)))
	}

	for _, ingredient := range ingredientsMass {
		documentBuilder.WriteString(fmt.Sprintf("\n- %s: %.0fg", ingredient.ingredientName, math.Round(ingredient.massValue)))
	}

	documentBuilder.WriteString(fmt.Sprintf("\n\n**Container**:\n- %s\n- %.0fcm x %.0fcm x %.0fcm\n- %.0fcm³",
		values.ContainerName, values.Length, values.Width, values.Depth, math.Round(containerVolume)))
	documentString := documentBuilder.String()
	return documentString

}
