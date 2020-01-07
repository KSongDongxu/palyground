#!/bin/bash 

src=./hello/
dst=./world/
 
argstr=`ls $src | grep -v '2.txt\|lo' | xargs`
args=($argstr)
for arg in "${args[@]}"
do
    rm -rf "${dst}/${arg}"
    cp -R "${src}/${arg}" "${dst}/${arg}"
    chown -R root:root "${dst}/${arg}"
done

