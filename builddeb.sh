#!/bin/bash

go build .

mkdir -p -m 0775 ./deb/DEBIAN

echo "Package: wallpaperloader
Version: $1
Provides: wallpaperloader
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Roman Covanyan <rs@tsov.pro>
Description: WallPaperLoader $1
 WallPaperLoader $1 daemon hourly load and write into file /var/wallpaperloader/wallpaper.jpg a random wallpaper from images.yandex.ru, Bing or Unsplash
" > ./deb/DEBIAN/control

chmod 0664 ./deb/DEBIAN/control

echo "/etc/systemd/system/wallpaperloader.service" > ./deb/DEBIAN/conffiles
chmod 0664 ./deb/DEBIAN/conffiles

echo "#!/bin/bash

systemctl daemon-reload
systemctl enable wallpaperloader.service
systemctl start wallpaperloader
" > ./deb/DEBIAN/postinst

chmod 0775 ./deb/DEBIAN/postinst

echo "#!/bin/bash

systemctl stop wallpaperloader
systemctl disable wallpaperloader
exit 0
" > ./deb/DEBIAN/prerm

chmod 0775 ./deb/DEBIAN/prerm

echo "#!/bin/bash

systemctl daemon-reload
" > ./deb/DEBIAN/postrm

chmod 0775 ./deb/DEBIAN/postrm

mkdir -p -m 0755 ./deb/usr/local/bin
mkdir -p -m 0755 ./deb/etc/systemd/system
#mkdir -m 0666 -p ./deb/usr/share/wall
cp ./wallpaperloader ./deb/usr/local/bin/
cp ./wallpaperloader.service ./deb/etc/systemd/system/

fakeroot dpkg-deb --build ./deb

mv ./deb.deb ./wallpaperloader_$1_amd64.deb
lintian --no-tag-display-limit ./wallpaperloader_$1_amd64.deb

rm -r ./deb