<?php
if (isset($_GET['lang'])) {
    $lang = $_GET['lang'];
    $file_path = "lang/$lang";
    
    // Check if the file exists
    if (file_exists($file_path)) {
        include($file_path);
    } else {
        echo "Language file not found.";
    }
} else {
    echo "No language selected.";
}
?>