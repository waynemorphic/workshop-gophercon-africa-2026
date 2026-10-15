# Design: Wild, Observed

Inspired by a Dribbble reel: one large rounded photo filling the frame, a big elegant serif
headline over it, and on hover the photo slowly zooms, the headline blurs, and frosted-glass
captions float up. Quiet, premium, photo-first. Night colours.

## Colours (`web/assets/css/app.css`, `:root`)

| Variable | Hex | Use |
|---|---|---|
| `--charcoal` | `#141414` | Page background, `theme-color` |
| `--slate` | `#2A2A2A` | Photo frames before load, tiles without photos, footer panel |
| `--bone` | `#EDEAE4` | Main lettering |
| `--muted` | `#A8A49C` | Secondary text, nav links |
| `--moss` | `#8FA98A` | Accent: focus rings, "View the record" links |
| `--frost` | `rgba(20,20,20,.55)` + 14px backdrop blur | Captions over photos |

## Lettering

- Headings: Cormorant Garamond 400/500 (`--font-display`), large, tight leading (0.92 on the banner).
- Body and labels: Inter 400/500 (`--font-body`). Labels and nav are small uppercase with wide letter-spacing.
- Both from Google Fonts in the layout's `<head>`.

## Shape and spacing

- Side gutter `--gutter`: 16–32px. Radii: banner and footer 28px, tiles 18px, captions 12px.
- Easing `--ease`: `cubic-bezier(0.22, 1, 0.36, 1)`, slow (0.5–1.4s) for zooms.

## Components

- **Top bar**: paw logo + "Wild, Observed" wordmark left, uppercase nav right (Sightings, About).
- **Banner**: the first sighting with a photo, cover-fit in a rounded frame, dark gradient bottom-left,
  headline "Wild, observed." bottom-left, "See the wall" outlined button top-right. Hover: photo
  zooms, headline blurs, a white dot and a frosted note ("X, seen in Y") appear. On touch screens the
  note always shows and nothing blurs.
- **The wall**: CSS columns photo wall; tiles cycle through 3:4, 1:1 and 4:3 shapes (fixed, so lazy
  images don't jump). Frosted caption with name, Latin name, place, date, count, observer and a
  record link: hidden until hover/focus on a mouse, always shown on touch.
- **Tiles without photos**: slate panel with the animal's name in italic serif.
- **Footer** (`#about`): slate panel with a serif line about GBIF.
- Section heading fades in with the reveal block; individual tiles don't.

## Avoid

- Bright or saturated colours; anything busy competing with the photos.
- Fading every tile in on scroll.
