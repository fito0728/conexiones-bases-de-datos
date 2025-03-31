<?php
$host = "localhost";
$user = "root"; // Usuario por defecto en XAMPP
$password = ""; // Por defecto está vacío
$database = "tu_base_de_datos";

$conn = new mysqli($host, $user, $password, $database);

if ($conn->connect_error) {
    die("Error de conexión: " . $conn->connect_error);
}

echo "Conexión exitosa a MySQL en XAMPP";
?>