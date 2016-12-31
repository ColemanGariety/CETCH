#!/bin/bash

src=$(mktemp).go
echo "$1" >> $src
echo $src

dist=$(mktemp).go
go build -o $dist $src
