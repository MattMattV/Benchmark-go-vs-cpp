#!/bin/bash

cpt=0
echo "" > resultsGoPlusCPP


echo "langage,nbCores,segments,temps\n" >> resultsGoPlusCPP

for i in 1 2 4 8 16 32
do
        cpt=0
        segments=0
        langage="GO"
        while [ $cpt -lt 10 ]
        do
                segments=$(($segments+100000000))

                execTime=$( /usr/bin/time --format %U 2>&1 ./CalculGoCanaux $i -100 100 $segments > /dev/null )
                echo "$langage,$i,$segments,$execTime" >> resultsGoPlusCPP
                cpt=$((cpt + 1))
        done

        echo -e "\n" >> resultsGoPlusCPP

        cpt=0
        segments=0
        langage="C+"
        while [ $cpt -lt 10 ]
        do
                segments=$(($segments+100000000))

                execTime=$( /usr/bin/time --format %U 2>&1 ./Calcul.cxx $i -100 100 $segments > /dev/null )
                echo "$langage,$i,$segments,$execTime" >> resultsGoPlusCPP
                cpt=$((cpt + 1))
        done
        echo -e "\n\n" >> resultsGoPlusCPP

done
