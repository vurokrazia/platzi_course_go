docker container run -dit --name go_curse  --mount type=bind,source="$(pwd)",target=/go/src/  go_practice:1 