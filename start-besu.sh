docker container rm besu-testnet;

# docker run --name besu-testnet \
#   --rm \
#   --privileged \
#  --publish 10000:22 \
#  --publish 10001:3000 \
#  --publish 10002:8545 \
#  --publish 10003:8546 \
#  --publish 10004:9001 \
#  --publish 10005:9081 \
#  --publish 10006:9082 \
#  --publish 10007:9083 \
#  --publish 10008:9090 \
#  --publish 10009:18545 \ 
#  --publish 10010:20000 \
#  --publish 10011:20001 \
#  --publish 10012:20002 \
#  --publish 10013:20003 \
#  --publish 10014:20004 \
#  --publish 10015:20005 \
#  --publish 10016:25000 besu-mp-testnet-1 ;

docker compose up besu;
