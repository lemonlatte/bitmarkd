#!/bin/bash

add-apt-repository -y ppa:longsleep/golang-backports
apt-get -q update && apt-get -yqq install golang git

add-apt-repository -y ppa:bitmark/bitmarkd-$BITMARKD_VERSION
apt-get -q update && apt-get -yqq install bitmarkd bitmark-cli bitmark-info

sed -ie '/add_port("0.0.0.0", 2135)/d' /etc/bitmarkd.conf.sub
sed -ie '/add_port("0.0.0.0", 2136)/d' /etc/bitmarkd.conf.sub
sed -ie '/add_port("0.0.0.0", 2138)/d' /etc/bitmarkd.conf.sub
sed -ie '/add_port("0.0.0.0", 2139)/d' /etc/bitmarkd.conf.sub
sed -ie '/announce_ips\ =\ interface_public_ips/s/^--//g' /etc/bitmarkd.conf

git clone https://github.com/bitmark-inc/bitmark-wallet
cd bitmark-wallet && git checkout v0.6.3 && mkdir bin
go build -o bin -ldflags "-X main.version=0.6.3" ./...
cd

mv bitmark-wallet/bin/* /usr/local/bin/

rm -rf go bitmarkd
apt-get -y purge golang git
