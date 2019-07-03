pkgname=wallpaperloader
_pkgname=wallpaperloader
pkgver=1.4
pkgrel=1
pkgdesc="wallpaperloader"
arch=('i686' 'x86_64')
license=('GPL')
depends=(
)

source=(
	"$_pkgname::https://github.com/covrom/wallpaperloader#branch=${BRANCH:-master}"
)

md5sums=(
	'SKIP'
)

backup=(
)

build() {
	echo ":: Building binary"
	go build -v ${srcdir}/../.
}

package() {
	install -DT "wallpaperloader" "$pkgdir/usr/bin/wallpaperloader"
}
