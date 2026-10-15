package pages

import (
	"fmt"
	"workshop/internal/domain/sighting"
	"workshop/web/views/components"
)

// heroPhoto returns the first sighting that has a photo, for the banner.
func heroPhoto(sightings []sighting.Sighting) (sighting.Sighting, sighting.Media, bool) {
	for _, s := range sightings {
		if m, ok := components.FirstPhoto(s); ok {
			return s, m, true
		}
	}
	return sighting.Sighting{}, sighting.Media{}, false
}

// summaryLine sums up the wall in one sentence.
func summaryLine(sum sighting.Summary) string {
	return fmt.Sprintf("%d sightings of %d species, across %d countries.", sum.Sightings, sum.Species, sum.Countries)
}
