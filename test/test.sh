#!/bin/bash

if [ -z "$1" ]; then
	export AVALAM_COMPILE_ENV="dev"
else
	export AVALAM_COMPILE_ENV=$1
fi

rm ./.DS_Store 2> /dev/null
rm */.DS_Store 2> /dev/null
rm */*/.DS_Store 2> /dev/null
rm */*/*/.DS_Store 2> /dev/null

version=$(cat config/compile_version_base)
suffix=$(cat config/compile_version_suffix)
suffix=$(($suffix+1))
log_compile_path="logs/compile.logs"

echo "Compilation de `$build/$AVALAM_COMPILE_ENV-$version.$suffix.exe`..."
echo "[$(date)] ++ `Compilation de build/$AVALAM_COMPILE_ENV-$version.$suffix.exe` ..." >> $log_compile_path

g++ src/main.cpp src/classes/*/*.class.cpp -o build/$AVALAM_COMPILE_ENV-$version.$suffix.exe 

if [ $? -eq 0 ]; then
	echo $suffix > config/compile_version_suffix
	echo "build/$AVALAM_COMPILE_ENV-$version.$suffix.exe" > config/last_version
	echo "Compilation réussie"
fi

git add $log_compile_path build/$AVALAM_COMPILE_ENV-$version.$suffix.exe > /dev/null 2> /dev/null
git commit -m "Compilation de build/dev-..exe ..." > /dev/null 2> /dev/null
git push > /dev/null 2> /dev/null &
