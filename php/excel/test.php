<?php

require "vendor/autoload.php";

use PhpOffice\PhpSpreadsheet\IOFactory;
use PhpOffice\PhpSpreadsheet\Worksheet\MemoryDrawing;

$imageFilePath = './uploads/imgs/'; //图片本地存储的路径
if (!file_exists($imageFilePath)) { //如果目录不存在则递归创建
    mkdir($imageFilePath, 0777, true);
}

try {
    $inputFileName = './files/1.xlsx'; //包含图片的Excel文件
    $objRead = IOFactory::createReader('Xlsx');
    $objSpreadsheet = $objRead->load($inputFileName);
    $objWorksheet = $objSpreadsheet->getSheet(0);
    // $data = $objWorksheet->toArray();

    $spreadsheet = $objSpreadsheet;
    $i = 0;

    // var_dump($spreadsheet->getActiveSheet()->getDrawingCollection());
    foreach ($spreadsheet->getActiveSheet()->getDrawingCollection() as $drawing) {
        var_dump($drawing->getPath());
        // // var_dump($drawing->getMimeType());
        // if ($drawing instanceof MemoryDrawing) {
        //     ob_start();
        //     call_user_func(
        //         $drawing->getRenderingFunction(),
        //         $drawing->getImageResource()
        //     );
        //     $imageContents = ob_get_contents();
        //     ob_end_clean();
            
        //     switch ($drawing->getMimeType()) {
        //         case MemoryDrawing::MIMETYPE_PNG:
        //             $extension = 'png';
        //             break;
        //         case MemoryDrawing::MIMETYPE_GIF:
        //             $extension = 'gif';
        //             break;
        //         case MemoryDrawing::MIMETYPE_JPEG:
        //             $extension = 'jpg';
        //             break;
        //     }
        // } else {
        //     if ($drawing->getPath()) {
        //         // Check if the source is a URL or a file path
        //         if ($drawing->getIsURL()) {
        //             $imageContents = file_get_contents($drawing->getPath());
        //             $filePath = tempnam(sys_get_temp_dir(), 'Drawing');
        //             file_put_contents($filePath, $imageContents);
        //             $mimeType = mime_content_type($filePath);
        //             // You could use the below to find the extension from mime type.
        //             // https://gist.github.com/alexcorvi/df8faecb59e86bee93411f6a7967df2c#gistcomment-2722664
        //             $extension = File::mime2ext($mimeType);
        //             unlink($filePath);
        //         } else {
        //             $zipReader = fopen($drawing->getPath(), 'r');
        //             $imageContents = '';
        //             while (!feof($zipReader)) {
        //                 $imageContents .= fread($zipReader, 1024);
        //             }
        //             fclose($zipReader);
        //             $extension = $drawing->getExtension();
        //         }
        //     }
        // }
        // $myFileName = '00_Image_' . ++$i . '.' . $extension;
        // file_put_contents($myFileName, $imageContents);
    }

    // foreach ($objWorksheet->getDrawingCollection() as $drawing) {
    //     list($startColumn, $startRow) = Coordinate::coordinateFromString($drawing->getCoordinates());
    //     $imageFileName = $drawing->getCoordinates() . mt_rand(1000, 9999);
    //     var_dump(111);
    //     switch ($drawing->getExtension()) {
    //         case 'jpg':
    //         case 'jpeg':
    //             $imageFileName .= '.jpg';
    //             $source = imagecreatefromjpeg($drawing->getPath());
    //             imagejpeg($source, $imageFilePath . $imageFileName);
    //             break;
    //         case 'gif':
    //             $imageFileName .= '.gif';
    //             $source = imagecreatefromgif($drawing->getPath());
    //             imagegif($source, $imageFilePath . $imageFileName);
    //             break;
    //         case 'png':
    //             $imageFileName .= '.png';
    //             $source = imagecreatefrompng($drawing->getPath());
    //             imagepng($source, $imageFilePath, $imageFileName);
    //             break;
    //     }
    //     $startColumn = ABC2decimal($startColumn);
    //     $data[$startRow-1][$startColumn] = $imageFilePath . $imageFileName;
    // }
    // var_dump($data);die();
} catch (\Exception $e) {
    throw $e;
}

function ABC2decimal($abc)
{
    $ten = 0;
    $len = strlen($abc);
    for ($i = 1; $i <= $len; $i++) {
        $char = substr($abc, 0 - $i, 1); //反向获取单个字符

        $int = ord($char);
        $ten += ($int - 65) * pow(26, $i - 1);
    }
    return $ten;
}
