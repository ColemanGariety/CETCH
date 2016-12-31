#!/bin/bash

src=$(mktemp).go
echo "$1" >> $src

dist=$(mktemp).go
go build -o $dist $src

mbox -s -n -c -R -i -- $dist
