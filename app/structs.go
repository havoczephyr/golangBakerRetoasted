package app

type ingredientPercent struct {
	ingredientName  string
	percentageValue float64
}

type flourPercent struct {
	flourName       string
	percentageValue float64
}

type ingredientMass struct {
	ingredientName string
	massValue      float64
}

type flourMass struct {
	flourName string
	massValue float64
}

type ImportValues struct {
	TempAtmosphere bool    `json:"warmWeather"`
	Length         float64 `json:"containerLength"`
	Width          float64 `json:"containerWidth"`
	Depth          float64 `json:"containerDepth"`
	ContainerName  string  `json:"containerName"`
}
