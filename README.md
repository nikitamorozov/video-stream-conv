## Before start
You have to change ```config.json``` to use this service

## Run project
```docker-compose up -d```

## Todo
* implement queue processing

## Links
To send video file, u need POST tp url
http://converter.club500.com/api/v1/convert
with ```file``` param

To get converted stream see on link below
ex: http://video.club500.com/ef763ee724e0cb62b2cfbb3d29a1234672ea8091_playlist.m3u8

## Todo
* JWT
* Separate worker
* Possibility to run few workers