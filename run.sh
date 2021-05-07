make -C server compile
yarn --cwd ui install
yarn --cwd ui build
mv ui/dist server/static
docker-compose -f deploy/docker-compose.yaml up
