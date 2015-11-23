#!/bin/bash

cpt=0
segments=10
echo "" > resultsCpp
while [ $cpt -lt 100 ]
do

	execTime=$( /usr/bin/time --format %U 2>&1 Calcul.cxx -100 100 $segments > /dev/null )
	echo "$segments,$execTime" >> resultsCpp
	cpt=$((cpt + 1))
	segments=$((segments * 2))
	echo "+1 tour"

done
