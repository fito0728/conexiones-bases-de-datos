#include <iostream>
#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>
#include <mongocxx/database.hpp>
#include <mongocxx/uri.hpp>

int main() {
    // Inicializa la instancia de MongoDB
    mongocxx::instance inst{}; 

    // Crea un cliente de MongoDB que se conecta a localhost:27017
    mongocxx::client conn{mongocxx::uri{"mongodb://localhost:27017"}};

    // Accede a una base de datos específica (en este caso 'test')
    mongocxx::database db = conn["test"];

    // Accede a una colección dentro de la base de datos
    mongocxx::collection coll = db["mi_coleccion"];

    // Crear un documento BSON para insertar en la colección
    bsoncxx::builder::stream::document document{};
    document << "nombre" << "Juan" << "edad" << 25;

    // Insertar el documento en la colección
    coll.insert_one(document.view());

    std::cout << "Documento insertado con éxito" << std::endl;

    return 0;
}