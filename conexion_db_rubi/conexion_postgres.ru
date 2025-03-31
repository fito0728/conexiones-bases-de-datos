require 'pg'

conn = PG.connect(
  dbname: 'tu_base_de_datos',
  user: 'tu_usuario',
  password: 'tu_contraseña',
  host: 'localhost',
  port: 5432
)

puts "Conexión exitosa a PostgreSQL"