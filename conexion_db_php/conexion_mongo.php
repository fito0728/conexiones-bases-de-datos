<?php
require 'vendor/autoload.php'; // Necesita Composer

$client = new MongoDB\Client("mongodb://localhost:27017");
$database = $client->tu_base_de_datos;
$collection = $database->tu_coleccion;

echo "Conectado a MongoDB correctamente";
?>