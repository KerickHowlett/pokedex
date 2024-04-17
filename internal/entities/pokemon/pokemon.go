package pokemon

// Pokemon houses the essential information of a given Pokemon.
//
// Fields:
//   - BaseExperience: The base experience of the Pokemon, which effects the catch rate.
//   - Name: The name of the Pokemon.
type Pokemon struct {
	// BaseExperience represents the base experience of the Pokemon, which effects the catch rate.
	BaseExperience int `json:"base_experience"`
	// Name represents the name of the Pokemon.
	Name string `json:"name"`
}
