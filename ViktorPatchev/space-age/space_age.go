package space

type Planet string

const earth = 31557600

func calculateAge(planet float64, ageInSeconds float64) float64 {
	return ageInSeconds / planet
}

func Age(ageInSeconds float64, planet Planet) float64 {
	Planets := map[Planet]float64{
		"Earth":   earth,
		"Mercury": earth * 0.2408467,
		"Venus":   earth * 0.61519726,
		"Mars":    earth * 1.8808158,
		"Jupiter": earth * 11.862615,
		"Saturn":  earth * 29.447498,
		"Uranus":  earth * 84.01684,
		"Neptune": earth * 164.79132,
	}
	return calculateAge(Planets[planet], ageInSeconds)
}
