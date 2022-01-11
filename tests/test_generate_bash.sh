#!/bin/bash

i=1
# 10 minute
while [ $i -lt 60 ]
do
  generateblock -l
  sleep 1
  i=`expr $i + 1`
done
