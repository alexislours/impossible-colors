![GitHub Workflow Status](https://img.shields.io/github/workflow/status/alexislours/impossible-colors/Build%20releases) 
![GitHub release (latest by date)](https://img.shields.io/github/v/release/alexislours/impossible-colors) 
![GitHub](https://img.shields.io/github/license/alexislours/impossible-colors)

# impossible-colors

A quick app to generate images with alternating color values, making it possible to create "[impossible colors](https://en.wikipedia.org/wiki/Impossible_color)".


## Usage

Build the project by cloning the repository and running `go build` inside it.

Generate the image by running `./impossible-colors <width> <height> <color1> <color2>`.

eg: using `./impossible-colors 400 400 ff0 00f`

![Example](images/400x400_ff0_00f.png)

Prebuilt binaries for Windows, Linux and macOS can be found [here](https://github.com/alexislours/impossible-colors/releases/latest).