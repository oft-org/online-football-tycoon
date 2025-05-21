/wait-for-it.sh db:5432 -- echo "Database is up"

if [ ! -f /app/migrations_done ]; then
  echo "Running database migrations..."
  ./migrate up

   touch /app/migrations_done
else
  echo "Migrations have already been applied."
fi

exec /app/app
