1.docker run -p 2379:2379 -p 2380:2380 --name myetcd  -v /etc/ssl/certs/:/etc/ssl/certs/ quay.io/coreos/etcd:v3.4.0

2. raft 算法
3.quorum模型

4.安装etcd
docker run -p 2379:2379 -p 2380:2380 --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data --name myetcd gcr.io/etcd-development/etcd:v3.5.1 /usr/local/bin/etcd --name s1 --data-dir /etcd-data --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380 --initial-cluster s1=http://0.0.0.0:2380 --initial-cluster-token tkn --initial-cluster-state new --log-level info --logger zap --log-outputs stderr
