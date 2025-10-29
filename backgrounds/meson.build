# Current wallpaper set

wallpapers = [
  'Twigonometry.jpg',
  'cefn-sidan.jpg',
  'matador-rocks.jpg',
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
