#!/bin/bash
echo "this is test success!"
mkdir /results/
i=1
while(( i <= 100 )) 
do
echo $i
let "i +=1"
sleep 1
done
echo "results" > /results/results.txt
echo "logs" > /results/logs.txt
