#! /usr/bin/python
import csv
import re

p = re.compile('is:[\s]*([0-9]*)[\s](?:.*[\s]){2}BenchmarkChan[-0-9]*[\s]*[0-9]*[\s]*([0-9]*)[\sa-zA-Z/]*([0-9]*)[\sa-zA-Z/]*([0-9]*)[\sa-zA-Z/]*([0-9]*)BenchmarkLock[-0-9]*[\s]*[0-9]*[\s]*([0-9]*)[\sa-zA-Z/]*([0-9]*)[\sa-zA-Z/]*([0-9]*)[a-zA-Z]*([0-9]*)')
with open('./output/chan.csv', 'wb') as chan, open('./output/lock.csv', 'wb') as lock, open('./output/bash.out', 'rb') as bash:
    fieldnames = ["readperc","ns_op","B_op","allocs_op","secs"]
    chanWriter = csv.writer(chan)
    lockWriter = csv.writer(lock)
    chanWriter.writerow(fieldnames)
    lockWriter.writerow(fieldnames)
    for perc in p.findall(bash.read()):
        print 1234
        chanWriter.writerow([perc[0]]+list(perc[1:4]))
        lockWriter.writerow([perc[0]]+list(perc[5:8]))
    
