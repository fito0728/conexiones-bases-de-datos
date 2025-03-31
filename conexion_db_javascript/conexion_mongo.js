const { MongoClient } = require("mongodb");

const uri = "mongodb://localhost:27017";
const client = new MongoClient(uri);

async function conectarMongo() {
  try {
    await client.connect();
    console.log("Conexión exitosa a MongoDB Compass");

    const db = client.db("tu_base_de_datos");
    console.log("Colecciones:", await db.listCollections().toArray());
  } catch (error) {
    console.error("Error de conexión:", error);
  } finally {
    await client.close();
  }
}

conectarMongo();