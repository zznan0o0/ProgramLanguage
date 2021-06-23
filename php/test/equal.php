<?php

if(0 == null){
    echo "1";
}
if(0 === null){
    echo "2";
}
if(null === null){
    echo "3";
}

if(0 != "0"){
    echo "4";
}
if(0 !== "0"){
    echo "5";
}

if([]){
    echo "6";
}