# Current wallpaper set

wallpapers = [
  'AerialDesert.jpg',
  'Athlone-Shannon-1920x1080.png',
  'BluePeaks.jpg',
  'BogFlower.jpg',
  'BrittasLake.jpg',
  'Celeste.jpg',
  'Crags.png',
  'Excl_Autumn_Leaf.jpg',
  'Excl_Autumn_Retreat.jpg',
  'Excl_Focal_Leaf.jpg',
  'Excl_Focal_Spring.jpg',
  'Excl_Lazy_Summer_Leaf.jpg',
  'Excl_Lonely_Railway.jpg',
  'Excl_Vintage_Bicycles.jpg',
  'forest-street.jpg',
  'HermiteCrab2560x1600.jpg',
  'IMG_0597-solus.jpg',
  'IMG_0604-solus.jpg',
  'IMG_0608-solus.jpg',
  'IMG_0609-solus.jpg',
  'IMG_9656.png',
  'IMG_9710.png',
  'LakeSideView.jpg',
  'Mama.jpg',
  'Path.jpg',
  'Pines.jpg',
  'PubSide.jpg',
  'ShannonClonmacnoise.jpg',
  'Snow-Topped-Trees.jpg',
  'SolusFresh.png',
  'TulipBee.jpg',
  'TullamoreGrandCanal.jpg',
  'Twigonometry.jpg',
  'a-colorful-chaos.jpg',
  'book.png',
  'camilo-beach.jpg',
  'cefn-sidan.jpg',
  'chalk.jpg',
  'cloud-covered-mountains.jpg',
  'dazzle.jpg',
  'harbwr-penbre.jpg',
  'matador-rocks.jpg',
  'river-and-mountains.jpg',
  'the-devils-bridge.jpg',
  'waves-of-blue.jpg',
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
