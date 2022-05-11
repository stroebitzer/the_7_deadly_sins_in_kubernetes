docker run -it --rm --name app-e ueber/7sins-dataloss-e:1.0.0
docker run -it --rm --name app-s ueber/7sins-dataloss-s:1.0.0

docker stop app-e
docker stop app-s

