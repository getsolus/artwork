# Solus Artwork

## Backgrounds

Solus Artwork contains many backgrounds graciously provided by the community for the purposes of inclusion into Solus. Please see the below listing for the appropriate credits and licensing information:

- Except when declared otherwise, the majority of the photo images in `backgrounds/` have been kindly provided by John McCormack for inclusion into Solus. We are incredibly grateful for the beautiful images he has provided.
- AerialDesert.jpg is provided by [NASA](https://unsplash.com/photos/whDrFMucHkc) and licensed under [Unsplash License](https://unsplash.com/license).
- BluePeaks.jpg is provided by [Fabrizio Conti](https://unsplash.com/photos/k6GpdsPJSZw) and licensed under [Unsplash License](https://unsplash.com/license).
- Celeste.jpg is a modified wallpaper, originally sourced from [Paweł Czerwiński](https://unsplash.com/photos/SoB70WFVWGU) and licensed under [Unsplash License](https://unsplash.com/license).
- LakeSideView.jpg is provided by [Jake Hills](https://unsplash.com/photos/mI02K_LxlfU) and licensed under [Unsplash License](https://unsplash.com/license).
- Peeks.jpg is a modified wallpaper from [Pexels](https://www.pexels.com/photo/green-pine-trees-covered-with-fogs-under-white-sky-during-daytime-167699/) and is licensed under the [Pexels License](https://www.pexels.com/photo-license/).
- SolusFresh.png, book.png, and chalk.jpeg were kindly provided by Alejandro Seoane and are available under the main package license, GPL-2.0.
- Snow-Topped-Trees.jpg is a modified wallpaper from [Pexels](https://www.pexels.com/photo/aerial-photography-of-trees-1438761/) and is licensed under the [Pexels License](https://www.pexels.com/photo-license/).
- HermiteCrab2560x1600.jpg is kindly provided by Dirk Hohndel and is available under the CC-BY-3.0 license: https://creativecommons.org/licenses/by/3.0/us/
- `Excl_*.jpg` photographs are copyright (© 2016-2019) of Marius Nester, and are available exclusively for the Solus project per prior agreement. These may only be distributed as part of the official Solus release medium. All rights reserved.
- `IMG_9656.png` and `IMG_9710.png` © 2018-2019 Reetta Piiroinen. These may only be distributed as part of the official Solus release medium.
- `a-colorful-chaos.jpg` is provided by [Joel Filipe](https://unsplash.com/photos/k8apfKm-Md4) and licensed under [Unsplash License](https://unsplash.com/license).
- `camilo-beach.jpg` is provided by [Claudio Schwarz](https://unsplash.com/photos/kxAaw2bO1Z8) and licensed under [Unsplash License](https://unsplash.com/license).
- `cefn-sidan.jpg` is provided by Joseph Riches @joebonrichie and licensed under [CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/).
- `cloud-covered-mountains.jpg` is provided by [Nick Saxby](https://unsplash.com/photos/HFbnIMpYbcc) and licensed under [Unsplash License](https://unsplash.com/license).
- `dazzle.jpg` is provided by [Jot](https://linktr.ee/mustang164) for inclusion into Solus.
- `harbwr-penbre.jpg` is provided by Joseph Riches @joebonrichie and licensed under [CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/).
- `matador-rocks.jpg` is provided by Mica Semrick and is licensed under [CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/).
- river-and-mountains.jpg is provided by [Tony Reid](https://unsplash.com/photos/UMJUr3st0AE) and is licensed under [Unsplash License](https://unsplash.com/license).
- `the-devils-bridge.jpg` is provided by A.Landgraf and is licensed under [CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/).
- forest-street is a modified wallpaper from [Pexels](https://www.pexels.com/photo/forest-pathways-photo-1996051/) and is licensed under the [Pexels License](https://www.pexels.com/photo-license/).
- `waves-of-blue.jpg` is a modified wallpaper from [Pexels](https://unsplash.com/photos/_nWaeTF6qo0) and is licensed under the [Pexels License](https://www.pexels.com/photo-license/).

## Branding and Logos 

- Copyright © 2015-2021 Solus Project.
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
