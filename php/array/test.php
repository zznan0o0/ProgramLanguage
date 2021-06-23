<?php

$a = '[4,5,6]';

var_dump(json_decode($a, true));

// eval('$arr = '.$a.';');

// function stringToArray($str){
//     var_dump(eval("return " . $str . ";"));
// }
// var_dump(stringToArray($a));