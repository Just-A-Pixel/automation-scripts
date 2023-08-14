#!/usr/bin/env bash

script_name="$(echo -e "${1}" | tr -d '[:space:]')"

echo "$script_name"

if [ ! -d "./scripts/$script_name" ]; then
	echo "creating $script_name"
  	mkdir -p "./scripts/$script_name";
else
	echo "$script_name pkg already exists"
	exit
fi

echo "package main

func main() {

}
" > ./scripts/"$script_name"/script.go