docker rmi -f diplom
docker rm -f $(docker ps -a -q)
echo "removing container"
pause
docker rmi $(docker image ls -a -q)
pause
docker image prune