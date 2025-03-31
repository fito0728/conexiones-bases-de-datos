const { Client } = require("pg");

const client = new Client({
  user: "tu_usuario",
  host: "localhost",
  database: "tu_base_de_datos",
  password: "tu_contraseña",
  port: 5432,
});

async function conectarPostgres() {
  try {
    await client.connect();
    console.log("Conexión exitosa a PostgreSQL");
  } catch (error) {
    console.error("Error de conexión:", error);
  } finally {
    await client.end();
  }
}

conectarPostgres();
