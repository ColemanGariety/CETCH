#!/bin/bash

src=$(mktemp).go
echo "$1" >> $src

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
            bash -c "TIMEFORMAT='%3R'; time python2 $src")

echo $out
