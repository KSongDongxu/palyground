#!/bin/bash 
# 拷贝文件到另外一个文件夹 
# 排除个别文件或文件夹

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

