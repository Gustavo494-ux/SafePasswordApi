docker compose down safepassword
docker compose build safepassword
clear
docker compose up safepassword -d
docker cp safepassword:/usr/src/app/SafePasswordApi/docs/ ./
# docker stop safepassword
docker ps