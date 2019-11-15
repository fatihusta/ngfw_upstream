-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Format: 3.0 (quilt)
Source: wireguard
Binary: wireguard, wireguard-dkms, wireguard-tools
Architecture: linux-any all
Version: 0.0.20191012-1
Maintainer: Daniel Kahn Gillmor <dkg@fifthhorseman.net>
Homepage: https://www.wireguard.com
Standards-Version: 4.4.1
Vcs-Browser: https://salsa.debian.org/debian/wireguard
Vcs-Git: https://salsa.debian.org/debian/wireguard.git
Testsuite: autopkgtest
Testsuite-Triggers: iproute2, wget
Build-Depends: debhelper-compat (= 12), dkms, libmnl-dev, pkg-config, systemd
Package-List:
 wireguard deb net optional arch=all
 wireguard-dkms deb kernel optional arch=all
 wireguard-tools deb net optional arch=linux-any
Checksums-Sha1:
 73c6f4308653805e0b2a7edaa5d9f3af34411317 331176 wireguard_0.0.20191012.orig.tar.xz
 681aa44dfaa5899991b027c1c6ad1b68fc67793f 25268 wireguard_0.0.20191012-1.debian.tar.xz
Checksums-Sha256:
 f98968ec83b334ffa8eddf401f896c2ab49da1d4b6bc1c02e2219fe589e271d1 331176 wireguard_0.0.20191012.orig.tar.xz
 31870b62bf9d04212c17cb83c1bd068c2796350a04e635fd273fba6cf3e74857 25268 wireguard_0.0.20191012-1.debian.tar.xz
Files:
 b6316792324437dc9cc6092d2f6f43a7 331176 wireguard_0.0.20191012.orig.tar.xz
 df9a76671ff7053b8c5267d915528633 25268 wireguard_0.0.20191012-1.debian.tar.xz

-----BEGIN PGP SIGNATURE-----

iHUEARYKAB0WIQTJDm02IAobkioVCed2GBllKa5f+AUCXaXu/AAKCRB2GBllKa5f
+JREAQCDwCwD8imLJn+heEcYIyaLeShkWSQuIF5+88Ks1pVaNAD/W8QH4xDqp/CO
glXO+owStGCVYtpgALtH23C6WCX3gQA=
=OmOG
-----END PGP SIGNATURE-----
