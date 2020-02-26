docker build -t nikita-m/stream-converter .
docker run --rm -d -p 6754:6754 -v $(pwd)/video:/var/app/video nikita-m/stream-converter
