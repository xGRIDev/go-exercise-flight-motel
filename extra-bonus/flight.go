// 07
package main

import "fmt"

type FlightDetails struct {
	ID          int
	NamePlane   string
	Airport     string
	Type        string
	Destination string
	Price       float64
}

func main() {
	flights := []FlightDetails{
		{
			ID:          01,
			NamePlane:   "Trans-Nusa",
			Airport:     "Soekarno-Hatta International Airport",
			Type:        "Economy",
			Destination: "To Singapore",
			Price:       1.166,
		},
		{
			ID:          02,
			NamePlane:   "Trans-Nusa",
			Airport:     "Soekarno-Hatta International Airport",
			Type:        "Economy",
			Destination: "To Singapore",
			Price:       1.166,
		},
		{
			ID:          03,
			NamePlane:   "Batik-Air",
			Airport:     "Soekarno-Hatta International Airport",
			Type:        "Economy",
			Destination: "To Singapore",
			Price:       1.166,
		},
		{
			ID:          04,
			NamePlane:   "Garuda-Indonesia",
			Airport:     "Soekarno-Hatta International Airport",
			Type:        "Economy",
			Destination: "To Singapore",
			Price:       2.166,
		},
	}
	fmt.Println(flights)
	for _, p := range flights {
		fmt.Printf("  - ID: %d, Plane: %s, Price: %.2f\n", p.ID, p.NamePlane, p.Price)
	}
	flights = append(flights, FlightDetails{NamePlane: "Batik-Air", ID: 3, Price: 12.00})
	fmt.Println("Products after adding third:")
	for _, p := range flights {
		fmt.Printf("  - ID: %d, Title: %s, Price: %.2f\n", p.ID, p.NamePlane, p.Price)
	}
}
