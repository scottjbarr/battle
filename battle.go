package battle

import (
	"fmt"
	"math/rand"
)

type Battle struct {
	Fighters []*Fighter
}

func NewBattle(maxFighters int64) *Battle {
	fighters := make([]*Fighter, maxFighters, maxFighters)

	for i := range fighters {
		f := RandomFighter()
		fighters[i] = &f
	}

	return &Battle{
		Fighters: fighters,
	}
}

func (b *Battle) Turn() {
	alive := map[int]bool{}

	for i, f := range b.Fighters {
		if f.IsAlive() {
			alive[i] = true
		}
	}

	var keys []int
	for k := range alive {
		keys = append(keys, k)
	}

	shuffle(keys)

	// get the attacker
	attacker := b.Fighters[keys[0]]

	// remove the first key
	keys = append(keys[:0], keys[1:]...)

	// get the target
	target := b.Fighters[keys[0]]

	damage := attacker.Hit(target)
	fmt.Printf("%s hit %s doing %v damage\n", attacker, target, damage)

	if !target.IsAlive() {
		fmt.Printf("    %v is dead! :(\n", target.Name)
		fmt.Printf("    %v/%v fighters\n", len(b.Alive()), len(b.Fighters))
	}
}

func (b Battle) Finished() bool {
	return len(b.Alive()) == 1
}

func (b Battle) Alive() []*Fighter {
	alive := []*Fighter{}

	for _, f := range b.Fighters {
		if f.IsAlive() {
			alive = append(alive, f)
		}
	}

	return alive
}

func (b Battle) Winner() *Fighter {
	alive := b.Alive()

	if len(alive) == 1 {
		return alive[0]
	}

	return nil
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
