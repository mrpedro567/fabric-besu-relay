docker run --name besu-testnet \
  --rm \
  --privileged \
  --publish 10000:22 \
  --publish 10001:3000 \
  --publish 10002:8545 \
  --publish 10003:8546 \
  --publish 10004:9001 \
  --publish 10005:9081 \
  --publish 10006:9082 \
  --publish 10007:9083 \
  --publish 10008:9090 \
  --publish 10009:18545 \
  --publish 10010:20000 \
  --publish 10011:20001 \
  --publish 10012:20002 \
  --publish 10013:20003 \
  --publish 10014:20004 \
  --publish 10015:20005 \
  --publish 10016:25000 besu-mp-testnet-1 ;

docker run \
	--privileged \
	--env FABRIC_VERSION=2.2.13 \
	--name fabric-testnet-1 \
	--publish 20000:22 \
	--publish 20001:2375\
	--publish 20002:2376\
	--publish 20003:5984\
	--publish 20004:6984\
	--publish 20005:7050\
	--publish 20006:7051\
	--publish 20007:7054\
	--publish 20008:7984\
	--publish 20009:8051\
	--publish 20010:8054\
	--publish 20011:8984\
	--publish 20012:9001\
	--publish 20013:9051\
	--publish 20014:10051\
	fabric-testnet;

echo "${bold}************************";
echo "Endpoints besu testnet n1";
echo "${bold}************************";

echo "JSON-RPC HTTP service endpoint : http://localhost:10002";
echo "JSON-RPC WebSocket service endpoint : ws://localhost:10003";
echo "Web block explorer address : http://localhost:10016/";
echo "Prometheus address : http://localhost:10008";
echo "http://localhost:10001/d/XE4V0WGZz/besu-overview?orgId=1&refresh=10s&from=now-30m&to=now&var-system=All";


echo "${bold}************************";
echo "Endpoints Fabric testnet n1";
echo "${bold}************************";
