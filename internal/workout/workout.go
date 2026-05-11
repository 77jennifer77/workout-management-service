package workout

import "fmt"

type Workout struct {
	Owner     string     `json:"owner"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	Equipment Equipment  `json:"equipment"`
	Exercises []Exercise `json:"exercises"`
}

type Equipment struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Sets        int    `json:"reps"`
	Time        int    `json:"time"`
}

// CreateWorkout creates a workout, save new workout to db
func (w *Workout) CreateWorkout() error {
	fmt.Println(w)
	return nil
}

// GetWorkout gets a workout from db
func (w *Workout) GetWorkout() error {
	fmt.Println(w)
	return nil
}
