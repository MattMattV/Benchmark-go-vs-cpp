#!/bin/bash

cpt=0
segments=10
echo "" > resultsGoWaitGroup
while [ $cpt -lt 100 ]
do
	execTime=$( /usr/bin/time --format %U 2>&1 CalculGoWaitGroup -100 100 $segments > /dev/null )
	echo "$segments,$execTime" >> resultsGoWaitGroup
	cpt=$((cpt + 1))
	segments=$((segments * 2))
	echo "+1 tour"
done
