docker container rm fabric-testnet-1;
docker volume prune;
# docker run \
#	--privileged \
#	--env FABRIC_VERSION=2.2.13 \
#	--name fabric-testnet-1 \
#	--publish 20000:22 \
#	--publish 20001:2375\
#	--publish 20002:2376\
#	--publish 20003:5984\
#	--publish 20004:6984\
#	--publish 20005:7050\
#	--publish 20006:7051\
#	--publish 20007:7054\
#	--publish 20008:7984\
#	--publish 20009:8051\
#	--publish 20010:8054\
#	--publish 20011:8984\
#	--publish 20012:9001\
#	--publish 20013:9051\
#	--publish 20014:10051\
#	fabric-testnet;

docker compose up fabric ;
