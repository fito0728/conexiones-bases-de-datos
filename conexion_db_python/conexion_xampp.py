import mysql.connector

conn = mysql.connector.connect(
    host="localhost",
    user="root",  # Cambia si tienes otro usuario
    password="",  # Deja vacío si no has puesto contraseña en XAMPP
    database="tu_base_de_datos"
)

print("Conexión exitosa a MySQL en XAMPP")

conn.close()