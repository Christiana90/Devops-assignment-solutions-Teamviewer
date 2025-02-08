package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
)

// Mass struct holds the density of a material
type Mass struct {
	Density float64
}

// MassVolume interface ensures that any shape must implement density and volume calculations
type MassVolume interface {
	density() float64
	volume(dimension float64) float64
}

// Sphere struct represents an aluminum sphere, embedding Mass for density
type Sphere struct {
	Mass
}

// Cube struct represents an iron cube, embedding Mass for density
type Cube struct {
	Mass
}

// density() returns the density of a material
func (m Mass) density() float64 {
	return m.Density
}

// volume() calculates the volume of a sphere using the formula (4/3) * π * r³
func (s Sphere) volume(radius float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

// volume() calculates the volume of a cube using the formula side³
func (c Cube) volume(side float64) float64 {
	return math.Pow(side, 3)
}

// Handler function processes HTTP requests and calculates mass
func Handler(massVolume MassVolume) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dimensionStr := r.URL.Query().Get("dimension")
		dimension, err := strconv.ParseFloat(dimensionStr, 64)
		if err != nil {
			http.Error(w, "Invalid or missing dimension parameter", http.StatusBadRequest)
			return
		}
		weight := massVolume.density() * massVolume.volume(dimension)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%.2f", math.Round(weight*100)/100)))
	}
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <port>")
	// 	os.Exit(1)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable is required")
		os.Exit(1)
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Invalid port number:", port)
		os.Exit(1)
	}

	aluminiumSphere := Sphere{Mass{Density: 2.710}}
	ironCube := Cube{Mass{Density: 7.874}}

	http.HandleFunc("/aluminium/sphere", Handler(aluminiumSphere))
	http.HandleFunc("/iron/cube", Handler(ironCube))

	fmt.Printf("Server running on port %d\n", portInt)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", portInt), nil); err != nil {
		panic(err)
	}
}