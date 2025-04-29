#!/bin/bash

# Esperar a que la base de datos esté lista
/wait-for-it.sh db:5432 -- echo "Database is up"

# Verificar si las migraciones han sido ejecutadas previamente
if [ ! -f /app/migrations_done ]; then
  echo "Running database migrations..."
  # Ejecuta tus migraciones aquí, por ejemplo, con una herramienta de migraciones como 'golang-migrate'
  ./migrate up

  # Crear un archivo de marcador que indica que las migraciones ya se han ejecutado
  touch /app/migrations_done
else
  echo "Migrations have already been applied."
fi

# Ejecutar la aplicación (API)
exec /app/app
