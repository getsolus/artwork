# Current wallpaper set

wallpapers = [
  'AS08-16-2593.jpg',
  'bath-time.jpg',
  'blue-wallpaper.jpg',
  'boats-on-the-river.jpg',
  'cefn-sidan.jpg',
  'circumpolar-trails-sunset.jpg',
  'dessert-at-dusk.jpg',
  'digital-art-background.jpg',
  'driftwood-throne.jpg',
  'exhibición-freestyle-gp.jpg',
  'green-glass-vase.jpg',
  'happy-turtle.jpg',
  'komodo-dragon.jpg',
  'lower-antelope-canyon.jpg',
  'luxor-nile.jpg',
  'matador-rocks.jpg',
  'milky-way.jpg',
  'mount-batok-java.jpg',
  'red-cherry-leaves.jpg',
  'rosa-kent.jpg',
  'sebkha-d-oran.jpg',
  'Twigonometry.jpg',
  'vents-du-sud.jpg',
  'white-and-multicolored.jpg',
  'wild-golden-eagle.jpg',
]

wallpaper_dir = join_paths(path_datadir, 'backgrounds', 'solus')

mogrify = find_program('mogrify')
script = find_program(join_paths(meson.project_source_root(), 'scripts', 'optimizeImage.sh'))

optimized_background_targets = []

foreach wallpaper : wallpapers
  optimized_background_targets += custom_target(
    wallpaper,
    input: wallpaper,
    output: '@BASENAME@.jxl',
    command: [script, '@INPUT@', '@OUTPUT@'],
    install: true,
    install_dir: wallpaper_dir,
  )
endforeach

# Allow configuring the two wallpaper XML files
cdata = configuration_data()
cdata.set('prefix', path_prefix)

# Primary Wallpaper Definition
gnome_background_dir = join_paths(path_datadir, 'gnome-background-properties')
configure_file(
  input: 'solus-wallpapers.xml.in',
  output: 'solus-wallpapers.xml',
  configuration: cdata,
  install_dir: gnome_background_dir,
)

# MATE Wallpaper definition
mate_background_dir = join_paths(path_datadir, 'mate-background-properties')
configure_file(
  input: 'solus-wallpapers.xml.in',
  output: 'solus-wallpapers-mate.xml',
  configuration: cdata,
  install_dir: mate_background_dir,
)
