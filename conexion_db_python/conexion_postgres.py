import psycopg2

conn = psycopg2.connect(
    dbname="tu_base_de_datos",
    user="tu_usuario",
    password="tu_contraseña",
    host="localhost",
    port="5432"
)

print("Conexión exitosa a PostgreSQL")

conn.close()