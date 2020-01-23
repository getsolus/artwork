#!/bin/bash
set -e

rm -rf build
meson --prefix /usr build
ninja dist -C build

# Bump in tandem with meson.build, run script once new tag is up.
VERSION="27"
TAR="artwork-${VERSION}.tar.xz"
mv build/meson-dist/$TAR .

# Automatically sign the tarball with GPG key of user running this script
gpg --armor --detach-sign $TAR
gpg --verify "${TAR}.asc"
