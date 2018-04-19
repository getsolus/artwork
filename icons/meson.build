icons_dir = join_paths(path_datadir, 'icons', 'hicolor')

# Install the places icons
places_icon_sizes = [
    '256x256',
    '512x512',
    'scalable',
]

foreach icon_size : ['256x256', '512x512', 'scalable']
    if icon_size == 'scalable'
        suffix = '.svg'
    else
        suffix = '.png'
    endif
    icons = [
        '@0@/distributor-logo-solus@1@'.format(icon_size, suffix),
        '@0@/start-here-solus@1@'.format(icon_size, suffix),
    ]
    target_dir = join_paths(icons_dir, icon_size, 'places')
    install_data(
        icons,
        install_dir: target_dir,
    )
endforeach


# Install the mimetypes icons
mime_icon_sizes = [
    '16x16',
    '24x24',
    '32x32',
    '48x48',
    '64x64',
    '96x96',
    '256x256',
    '512x512',
    'scalable',
]

foreach icon_size : mime_icon_sizes
    if icon_size == 'scalable'
        suffix = '.svg'
    else
        suffix = '.png'
    endif
    icon_file = '@0@/application-x-solus-package@1@'.format(icon_size, suffix)
    target_dir = join_paths(icons_dir, icon_size, 'mimetypes')
    install_data(
        icon_file,
        install_dir: target_dir,
    )
endforeach