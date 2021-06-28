<?php

$a = '[4,5,6]';

// var_dump(json_decode($a, true));


function stringToArray($str){
    return eval("return " . $str . ";");
}
// var_dump(stringToArray($a));




$a1 = [1,2,3,4,5,6];
$a2 = [5,4,3,"2"];

$ia = array_intersect($a1, $a2);

var_dump($ia);


$ma = array_merge($a1, $a2);
$ma = array_values(array_unique($ma));
// var_dump($ma);

