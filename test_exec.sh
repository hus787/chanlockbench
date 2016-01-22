#! /bin/sh



for i in $(seq 100 $END); do
    export READ_PERCENTAGE=$i
    go test -bench . -benchmem
done
