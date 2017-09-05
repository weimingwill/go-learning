#!/bin/bash

#set -x

LOAD_TESTER_RC=load-tester
LOAD_TESTER_NAMESPACE=load-tester
LOAD_TESTER_LEADER_KEY="HostName"

ETCD_ELB=internal-tokyo-main-2-elb-etcd-1309004264.ap-northeast-1.elb.amazonaws.com
ETCD_PORT=4001
K8S_ELB=internal-tokyo-main-2-elb-k8s-839152653.ap-northeast-1.elb.amazonaws.com
K8S_PORT=8080

COUNTER=1

#ETCD CHANGE

LOAD_TESTER_LEADER_POD=`curl http://${ETCD_ELB}:${ETCD_PORT}/v2/keys/${LOAD_TESTER_LEADER_KEY} | jq .node.value`
COUNTER=1
LOAD_TESTER_LEADER_POD=${LOAD_TESTER_LEADER_POD//\"/}
echo "------- New leader node: $LOAD_TESTER_LEADER_POD -------------------"

POD_SELECTOR="name=load-tester"
POD_LIST=($(kubectl -s http://${K8S_ELB}:${K8S_PORT} get pods --selector=${POD_SELECTOR} --namespace=${LOAD_TESTER_NAMESPACE} | awk 'NR > 1 {print $1}'))

echo "------- Remove leader-node label for current pods-------------------"
for POD in "${POD_LIST[@]}"
do
    echo $POD
    kubectl -s http://${K8S_ELB}:${K8S_PORT} label pods ${POD} leader-node- --namespace=${LOAD_TESTER_NAMESPACE}
done

echo "-------- Label leader-node for new leader pod----------------------"
kubectl -s http://${K8S_ELB}:${K8S_PORT} label pods $LOAD_TESTER_LEADER_POD leader-node="true" --namespace=${LOAD_TESTER_NAMESPACE}