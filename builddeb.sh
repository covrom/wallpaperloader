#!/bin/bash

go build .

chmod -R 0755 ./deb

echo "Package: wallpaperloader
Version: $1
Provides: wallpaperloader
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Roman Covanyan <rs@tsov.pro>
Description: WallPaperLoader $1
 WallPaperLoader $1 daemon hourly load and write into file /usr/share/wall/wall.jpg a random wallpaper from images.yandex.ru, Bing or Unsplash
" > ./deb/DEBIAN/control

chmod 0644 ./deb/DEBIAN/control

echo "#!/bin/bash

mkdir -m 0666 /usr/share/wall
systemctl daemon-reload
systemctl enable wallpaperloader.service
systemctl start wallpaperloader
" > ./deb/DEBIAN/postinst

chmod 0755 ./deb/DEBIAN/postinst

echo "#!/bin/bash

systemctl stop wallpaperloader
systemctl disable wallpaperloader
exit 0
" > ./deb/DEBIAN/prerm

chmod 0755 ./deb/DEBIAN/prerm

echo "#!/bin/bash

systemctl daemon-reload
" > ./deb/DEBIAN/postrm

chmod 0755 ./deb/DEBIAN/postrm

mkdir -p ./deb/usr/local/bin
mkdir -p ./deb/usr/lib/systemd/system
cp ./wallpaperloader ./deb/usr/local/bin/
cp ./systemd/wallpaperloader.service ./deb/usr/lib/systemd/system/

fakeroot dpkg-deb --build ./deb

mv ./deb.deb ./wallpaperloader_$1_amd64.deb
lintian --no-tag-display-limit ./wallpaperloader_$1_amd64.deb