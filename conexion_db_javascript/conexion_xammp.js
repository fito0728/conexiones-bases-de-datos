const mysql = require("mysql2");

const connection = mysql.createConnection({
  host: "localhost",
  user: "root", // Cambia si tienes otro usuario
  password: "", // Deja vacío si no pusiste contraseña en XAMPP
  database: "tu_base_de_datos",
});

connection.connect((err) => {
  if (err) {
    console.error("Error de conexión:", err);
  } else {
    console.log("Conexión exitosa a MySQL en XAMPP");
  }
  connection.end();
});