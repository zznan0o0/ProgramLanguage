<?php
$t1=microtime(true);
$number = 0;

for($i=0; $i<100000000; $i++){
    $number += $i;
}

$t2=microtime(true);

echo $t2 - $t1;