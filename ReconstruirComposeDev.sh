docker compose down dev
docker compose build dev
clear
docker compose up dev -d
docker cp dev:/go/src/docs/ ./
# docker stop safepassword
docker ps