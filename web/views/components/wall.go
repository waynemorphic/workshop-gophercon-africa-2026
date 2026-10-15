package components

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"workshop/internal/domain/sighting"
	"workshop/internal/domain/species"
)

// DisplayName prefers the name a reader would recognise.
func DisplayName(sp species.Species) string {
	return cmp.Or(sp.VernacularName, sp.CanonicalName, sp.ScientificName, "Unknown species")
}

// Place joins whichever parts of the location GBIF supplied.
func Place(l sighting.Location) string {
	parts := slices.DeleteFunc(
		[]string{l.Locality, l.StateProvince, l.Country},
		func(s string) bool { return s == "" },
	)
	return cmp.Or(strings.Join(parts, ", "), "an unknown place")
}

// FirstPhoto returns the sighting's first still image, if it has one.
func FirstPhoto(s sighting.Sighting) (sighting.Media, bool) {
	for _, m := range s.Media {
		if m.Type == "StillImage" && m.URL != "" {
			return m, true
		}
	}
	return sighting.Media{}, false
}

// PhotoAlt describes a photo, falling back to the animal's name.
func PhotoAlt(s sighting.Sighting, m sighting.Media) string {
	return cmp.Or(m.Title, m.Description, "A "+DisplayName(s.Species)+" in the wild")
}

// tileShape cycles tile heights so the wall packs like a gallery.
func tileShape(i int) string {
	return [...]string{"tile--tall", "tile--square", "tile--wide", "tile--square", "tile--tall"}[i%5]
}

// individuals renders how many animals were seen in a single sighting.
func individuals(n int) string {
	if n == 1 {
		return "1 individual"
	}
	return fmt.Sprintf("%d individuals", n)
}
