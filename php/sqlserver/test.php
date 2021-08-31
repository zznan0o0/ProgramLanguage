<?php


$conn = sqlsrv_connect('118.31.21.122', array('Database' => 'AIS20210508222311', 'UID' => 'sa' , 'PWD' => 'Atjubo123456'));

if( $conn == false){
　　var_dump(sqlsrv_errors());exit;
}

$sql = "SELECT * FROM [dbo].[aaaa]";

$result = @sqlsrv_query($conn, $sql);
$re = @sqlsrv_fetch_array($result);
var_dump($re);