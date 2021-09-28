<?php

function c($arr){
    // foreach($arr as &$v2){
    //     $v2["a"] = 2;
    // }
    $arr = [["a" => 4]];
    var_dump($arr);
}


function test(){
    $arr = [["a" => 1]];
    foreach($arr as &$v){
        $v["a"] = 0;
    }

    // $arr2 = &$arr;

    // $arr2[0]["a"] = 3;

    c($arr);
    var_dump($arr);
}


test();