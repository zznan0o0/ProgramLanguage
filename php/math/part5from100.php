<?php


$p = 5;
$money = 100;
$surMoney = $money;

$moneyPart = [];

for($i = 0; $i < $p; $i++){
    $min = 1;
    $max = $surMoney - ($p - $i - 1) * $min;
    if($i == $p-1){
        $moneyPart[] = $surMoney;
        break;
    }
    $randMoney = rand($min, $max);
    $moneyPart[] = $randMoney;
    $surMoney -= $randMoney;
}

var_dump($moneyPart);