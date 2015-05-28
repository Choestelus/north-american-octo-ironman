#!/bin/bash
rm -f main
if [ "$#" -gt 0 ]; then
    echo $#
fi
for i in $(seq 1 $#);
do
    echo ${!i}
    rm ${!i}
done

