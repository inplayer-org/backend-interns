package house

import (
	"strings"
)

func Verse(v int) string {
	rhyme := "This is the "
	end := verses[0].object + " that " + verses[0].action + "."

	if v == 1 {
		return rhyme + end
	}

	rhyme += verses[v-1].object
	for ; v > 1; v-- {
		rhyme += "\nthat " + verses[v-1].action + " the "

		if v-1 == 1 {
			rhyme += end
		} else {
			rhyme += verses[v-2].object
		}
	}
	return rhyme
}

func Song() string {
	allverses := []string{}

	for v := 1; v <= len(verses); v++ {
		allverses = append(allverses, Verse(v))
	}
	return strings.Join(allverses, "\n\n")
}

var verses = []verse{
	verse{object: "house", action: "Jack built"},
	verse{object: "malt", action: "lay in"},
	verse{object: "rat", action: "ate"},
	verse{object: "cat", action: "killed"},
	verse{object: "dog", action: "worried"},
	verse{object: "cow with the crumpled horn", action: "tossed"},
	verse{object: "maiden all forlorn", action: "milked"},
	verse{object: "man all tattered and torn", action: "kissed"},
	verse{object: "priest all shaven and shorn", action: "married"},
	verse{object: "rooster that crowed in the morn", action: "woke"},
	verse{object: "farmer sowing his corn", action: "kept"},
	verse{object: "horse and the hound and the horn", action: "belonged to"},
}

type verse struct {
	object string
	action string
}
