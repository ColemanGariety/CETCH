#!/bin/bash

src=$(mktemp).hs
echo "$1" >> $src

dist=$(mktemp)
ghc -hidir /tmp -odir /tmp -o $dist $src

out=$(bwrap --ro-bind /usr /usr \
            --ro-bind /tmp /tmp \
            --proc /proc \
            --dev /dev \
            --unshare-net \
            --unshare-pid \
            --symlink usr/lib /lib \
            --symlink usr/lib64 /lib64 \
            --symlink usr/bin /bin \
            --symlink usr/sbin /sbin \
            bash -c "TIMEFORMAT='%3R'; time $dist")
echo $out
