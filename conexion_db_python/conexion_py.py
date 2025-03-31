from pymongo import MongoClient

# Conectar a MongoDB (asegúrate de que MongoDB esté en ejecución)
client = MongoClient("mongodb://localhost:27017/")

# Crear o conectar a una base de datos
db = client["mi_base_de_datos"]

# Crear o conectar a una colección
coleccion = db["mi_coleccion"]

# Insertar un documento en la colección
documento = {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}
coleccion.insert_one(documento)

print("Documento insertado correctamente")