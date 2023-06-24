package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Planet type
type Planet struct {
	Name      string
	Size      int
	Resources map[string]int
}

// generatePlanets generates planets and adds them to a slice
func generatePlanets(count int) []*Planet {
	planets := make([]*Planet, count)

	for i := 0; i < count; i++ {
		planets[i] = generateRandomPlanet()
	}

	return planets
}

// generateRandomPlanet creates a random planet
func generateRandomPlanet() *Planet {
	names := []string{"Earth", "Mars", "Jupiter", "Saturn", "Neptune", "Venus"}
	resources := map[string]int{"Fuel": rand.Intn(1000), "Oxygen": rand.Intn(1000)}

	return &Planet{
		Name:      names[rand.Intn(len(names))],
		Size:      rand.Intn(1000),
		Resources: resources,
	}
}

// planetProtection protects planets from being exploited or destroyed
func planetProtection(planets []*Planet) {
	for _, planet := range planets {
		// monitor planet resources
		go monitorResources(planet)

		// process resource extraction requests
		go processResourceExtractionRequests(planet)
	}
}

// monitorResources continuously monitors a planets resources
func monitorResources(planet *Planet) {
	for {
		fmt.Printf("Checking %s for resource depletion\n", planet.Name)
		time.Sleep(5 * time.Second)
		if isDepleted(planet.Resources) {
			fmt.Printf("%s's resources have been depleted, regenerating resources\n", planet.Name)
			generateResources(planet.Resources)
		}
	}
}

// isDepleted checks if a planets resources are empty
func isDepleted(resources map[string]int) bool {
	for _, v := range resources {
		if v > 0 {
			return false
		}
	}
	return true
}

// generateResources generates resources for a given planet
func generateResources(resources map[string]int) {
	for k, _ := range resources {
		resources[k] = rand.Intn(1000)
	}
}

// processResourceExtractionRequests processes requests for extracting resources from a planet
func processResourceExtractionRequests(planet *Planet) {
	for {
		fmt.Printf("Checking %s for resource extraction requests\n", planet.Name)
		time.Sleep(10 * time.Second)
		if request := getExtractionRequest(planet.Name); request.Quantity > 0 {
			fmt.Printf("Extracting %d %s from %s\n", request.Quantity, request.Resource, planet.Name)
			if planet.Resources[request.Resource] > request.Quantity {
				planet.Resources[request.Resource] -= request.Quantity
				fmt.Printf("Successfully extracted %d %s from %s\n", request.Quantity, request.Resource, planet.Name)
			} else {
				fmt.Printf("Not enough %s on %s to complete extraction\n", request.Resource, planet.Name)
			}
		}
	}
}

// ExtractionRequest type
type ExtractionRequest struct {
	Planet   string
	Resource string
	Quantity int
}

// getExtractionRequest generates a resource extraction request
func getExtractionRequest(planet string) ExtractionRequest {
	return ExtractionRequest{
		Planet:   planet,
		Resource: "Fuel",
		Quantity: rand.Intn(1000),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	planets := generatePlanets(5)
	planetProtection(planets)
}