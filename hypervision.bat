IF "%1" == "" GOTO end

docker build --build-arg="web_app_port=3000" -t gowebappimage .
echo GO_WEB_APP=3000 > .env
docker compose up

IF NOT "%2" == "op" GOTO end
start brave http://localhost:%1

:end