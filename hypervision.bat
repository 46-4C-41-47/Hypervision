IF "%1" == "" GOTO end

set goWebAppImagePort=%1

docker build -t gowebappimage .
docker compose up
start brave http://localhost:%1

:end