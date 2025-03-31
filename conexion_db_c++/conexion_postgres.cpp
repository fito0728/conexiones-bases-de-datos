#include <iostream>
#include <pqxx/pqxx>

int main() {
    try {
        pqxx::connection conn("dbname=mi_base user=mi_usuario password=mi_contraseña host=localhost port=5432");

        if (conn.is_open()) {
            std::cout << "Conectado a PostgreSQL" << std::endl;
        } else {
            std::cout << "Error en la conexión" << std::endl;
        }

        conn.disconnect();
    } catch (const std::exception &e) {
        std::cerr << e.what() << std::endl;
    }

    return 0;
}