package sub1

import (
	"github.com/rb-go/namegen"
	"math/rand"
)

func GetRandom() int {
	return rand.Intn(100)
}

func GetName() string {
	return namegen.GetName(0)
}
