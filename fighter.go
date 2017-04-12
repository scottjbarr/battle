package battle

import (
	"fmt"
	"math/rand"

	randomdata "github.com/Pallinder/go-randomdata"
)

type Fighter struct {
	Name        string
	Health      int64
	Power       int64
	Luck        float64
	Victories   int64
	TotalDamage int64
}

func RandomFighter() Fighter {
	f := Fighter{}

	f.Name = randomdata.SillyName()
	f.Health = random(100, 1000)
	f.Power = random(100, 200)
	f.Luck = rand.Float64()

	return f
}

func (f Fighter) IsAlive() bool {
	return f.Health > 0
}

func (f *Fighter) Hit(o *Fighter) int64 {
	modifier := 1.0

	r := rand.Float64()

	if r > f.Luck {
		modifier = modifier * r
	}

	damage := int64(float64(f.Power) * modifier)

	o.Health = o.Health - damage

	if !o.IsAlive() {
		f.Victories += 1
	}
	f.TotalDamage += damage

	return damage
}

func (f Fighter) String() string {
	return fmt.Sprintf("%v (H:%v P:%v L:%0.2f V:%v TD:%v)", f.Name, f.Health, f.Power, f.Luck, f.Victories, f.TotalDamage)
}

func random(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}
