<?php

$arr1 = [1,2,3];
$arr2 = [4,5,6];
$arr3 = [7,8,9];

var_dump(array_merge(...[$arr1, $arr2, $arr3]));

if([]){
    echo "arr";
}

echo "--------------";

$arr1 = [1,2,3,4,5];
$arr2 = [1,2,3,6,7];

$arr3 = ['0'=>1,'1'=>2,'2'=>3,'3'=>4,'4'=>5];
$arr4 = ['0'=>1,'1'=>2,'2'=>3,'3'=>6,'4'=>7];

$arr5 = ['0'=>1,'a'=>2,'b'=>3,'c'=>4,'4'=>5];
$arr6 = ['0'=>1,'a'=>2,'c'=>3,'d'=>6,'4'=>7];
var_dump(array_merge($arr1, $arr2));
var_dump($arr1+$arr2);
var_dump(array_keys(array_flip($arr1)+array_flip($arr2)));
echo "\n";
var_dump(array_merge($arr3, $arr4));
var_dump($arr3+$arr4);
var_dump(array_keys(array_flip($arr3)+array_flip($arr4)));
echo "\n";
var_dump(array_merge($arr5, $arr6));
var_dump($arr5+$arr6);



var_dump([1,2,3]+[4,5,3]);