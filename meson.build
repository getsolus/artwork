project(
    'artwork',
    ['c'],
    version: '30',
    license: [
        'CC-BY-3.0 ',
        'CC-BY-SA-4.0',
        'GPL-2.0-only',
    ],
    default_options: [
        'prefix=/usr',
        'sysconfdir=/etc',
    ],
)

path_prefix = get_option('prefix')
path_datadir = join_paths(path_prefix, get_option('datadir'))
path_licensedir = join_paths(path_datadir, 'licenses', meson.project_name())

subdir('backgrounds')
subdir('icons')

report = [
    '    Build configuration:',
    '    ====================',
    '',
    '    prefix:                                 @0@'.format(path_prefix),
    '    datadir:                                @0@'.format(path_datadir),
]

# Output some stuff to validate the build config
message('\n\n\n' + '\n'.join(report) + '\n\n')

# Install licenses and attribution
license_files = [
    'Attribution',
    'GPL-2.0-only',
    'CCBY-3.0',
    'CCBY-SA-4.0',
]

install_data(license_files,
    install_dir: path_licensedir
)
