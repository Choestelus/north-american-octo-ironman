#!/bin/bash
rm -f main
if [ "$#" -gt 0 ]; then
    echo $#
fi
for i in $(seq 1 $#);
do
    echo "removing [${!i}]"
    rm ${!i}
done

rm north-american-octo-ironman
rm -rf test/
