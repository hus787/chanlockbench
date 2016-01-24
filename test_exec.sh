#! /bin/sh

mkdir output
for i in $(seq 100 $END); do
    export READ_PERCENTAGE=$i
    go test -bench . -benchmem >>./output/bash.out 2>&1
    echo Done with $READ_PERCENTAGE
done
python result.py
