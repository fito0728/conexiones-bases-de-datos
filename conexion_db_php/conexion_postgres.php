<?php
$host = "localhost";
$port = "5432";
$dbname = "tu_base_de_datos";
$user = "tu_usuario";
$password = "tu_contraseña";

$conn = pg_connect("host=$host port=$port dbname=$dbname user=$user password=$password");

if ($conn) {
    echo "Conexión exitosa a PostgreSQL";
} else {
    echo "Error en la conexión";
}
?>