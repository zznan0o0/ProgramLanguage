<?php

echo strtotime("2021-01-01");
echo "\n";
echo strtotime("2021/01/01");
echo "\n";

var_dump(date("Y-m-d H:i:s", strtotime("2021-01-01")));