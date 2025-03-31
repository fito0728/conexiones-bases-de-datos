require 'mysql2'

client = Mysql2::Client.new(
  host: 'localhost',
  username: 'root',  # Cambia si tienes otro usuario
  password: '',  # Deja vacío si no has puesto contraseña en XAMPP
  database: 'tu_base_de_datos'
)

puts "Conexión exitosa a MySQL en XAMPP"