# Olasubomi Williams — Portfolio

A single-page HTMX portfolio powered by Go.

## Quick Start

1. **Copy your images** into `static/images/`:
   ```
   static/images/
   ├── 1000437712.jpg
   ├── Certificate Design1.png
   ├── back design black.png
   ├── back design regex.png
   ├── front top back bottom.png
   ├── grand robota.png
   ├── robota grand.png
   ├── black tshirt.png
   ├── tote design.jpg
   ├── Untitled design (1).png
   ├── Home.png
   ├── Untitled design.png
   ├── april designs (1).png
   ├── new week 10 march.png
   ├── New Week Flyers.png
   └── april designs.png
   ```

2. **Run the server:**
   ```bash
   go run main.go
   ```

3. **Open** [http://localhost:8080](http://localhost:8080)

## How it works

- Standard library `net/http` serves the site.
- `html/template` renders partials for HTMX or full pages for direct visits.
- `HX-Request` header detection swaps only `#main` content on nav clicks.
- Static assets (your images) are served from `/static/images/`.

## Project Structure

```
portfolio/
├── main.go              # Go server
├── go.mod               # Go module
├── templates/
│   ├── base.html        # Shell layout
│   ├── about.html
│   ├── software.html
│   ├── projects.html
│   ├── experience.html
│   └── contact.html
└── static/
    └── images/          # Your image assets
```
