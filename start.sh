docker build -t nikita-m/stream-converter .
docker run --rm -p 6754:6754 -v $(pwd)/video:/var/app/video nikita-m/stream-converter
