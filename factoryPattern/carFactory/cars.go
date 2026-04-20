package carfactory

// Car is the abstraction that all car types must implement.
// Client will depend ONLY on this.
type Car interface {
	GetName() string
}


// Now define different implementations.
type Sedan struct{}

func (s *Sedan) GetName() string {
	return "Sedan - Honda City"
}

type SUV struct{}

func (s *SUV) GetName() string {
	return "SUV - Mahindra XUV700"
}

type Hatchback struct{}

func (h *Hatchback) GetName() string {
	return "Hatchback - Maruti Swift"
}

func GetCar(carType string) Car {

	switch carType {
	case "sedan":
		return &Sedan{}
	case "suv":
		return &SUV{}
	case "hatchback":
		return &Hatchback{}
	default:
		return nil
	}
}