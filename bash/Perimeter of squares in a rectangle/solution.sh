#!/bin/bash
perimeter() {
  a=1 
  b=1     
  for (( i=0; i<$1+1; i++ )) 
  do
    let array+=$a
    fn=$((a + b)) 
    a=$b 
    b=$fn 
  done

  echo "$array * 4" | bc
}
perimeter $1