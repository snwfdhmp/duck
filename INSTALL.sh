#!/bin/bash

name="duck_install_$RANDOM"

mkdir $name
cd ./$name

echo -e "\033[1;33mDownloading latest sources ..."
wget https://raw.githubusercontent.com/snwfdhmp/duck/master/build/latest.tar.gz --output-document ./duck.tar > /dev/null 2> /dev/null
echo "Extracting ..."
tar -xzf ./duck.tar
echo "Building sources ..."
go build -o duck duck_install/cmd/duck/duck.go
echo "Installing duck ..."
sudo mv duck /usr/local/bin/duck
echo "Installing default configuration file ..."
sudo cp duck_install/ressources/duck.conf /etc/duck.conf
echo "Creating alias '@' for 'duck' ..."
sudo ln -s /usr/local/bin/duck /usr/local/bin/@ 2> /dev/null
echo "Cleaning ..."
cd ..
rm -rf ./$name
echo -e "Done ! Try '\033[1;36mduck version\033[1;33m' or '\033[1;31m@\033[1;36m version\033[1;33m'"
echo -e "\033[1;32mHave fun using duck\033[0;0m"