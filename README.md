# Solus Artwork

## Backgrounds

Solus Artwork contains many backgrounds graciously provided by the community for the purposes of inclusion into Solus. Please see `license/` for licensing details of each background.

## Branding and Logos

- Copyright © 2015-2025 Solus Project.
- The contents of `icons/` and the Solus logo itself is Copyright © 2016-2019 Solus Project. All Rights Reserved.

## Tooling

All source code for our wallpaper generators tool, available via `tooling/`, is distributed under GPL-2.0-only.

### Compiling

To compile our wallpaper generators tool, you must have [Go](https://golang.org) installed. If you are using Solus, you can do so by running:

``` bash
sudo eopkg install golang
```

Run `make build` to compile the tool, which will be put into `tooling/generate-wallpaper`

### Usage

To generate a new or updated `solus-wallpapers.xml.in` file, simply run: `make gen`
