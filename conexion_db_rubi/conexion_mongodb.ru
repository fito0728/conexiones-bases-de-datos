require 'mongo'

client = Mongo::Client.new('mongodb://localhost:27017/tu_base_de_datos')

puts "Conexión exitosa a MongoDB Compass"

# Listar colecciones como prueba
puts client.database.collection_names
