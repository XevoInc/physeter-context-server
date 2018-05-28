#/bin/sh

source .env

set -ex

: ${CLUSTER_NAME:="physeter-context-server"}
: ${SERVICE_NAME:?"Need to set SERVICE_NAME"}
: ${FAMILY_NAME:?"Need to set FAMILY_NAME"}
: ${AWS_ACCESS_KEY_ID:?"Need to set AWS_ACCESS_KEY_ID"}
: ${AWS_SECRET_ACCESS_KEY:?"Need to set AWS_SECRET_ACCESS_KEY"}

echo " # update task definition of $FAMILY_NAME"
DEF_FILE=`dirname $0`/${FAMILY_NAME}.json
DEFINITION=`cat $DEF_FILE`
aws ecs register-task-definition \
  --family $FAMILY_NAME \
  --network-mode awsvpc \
  --requires-compatibilities FARGATE \
  --cpu 512 \
  --memory 1024 \
  --task-role-arn "arn:aws:iam::059047002866:role/ecsTaskExecutionRole" \
  --execution-role-arn "arn:aws:iam::059047002866:role/ecsTaskExecutionRole" \
  --container-definitions "$DEFINITION"

echo " # update service of $SERVICE_NAME"
aws ecs update-service --cluster $CLUSTER_NAME --service $SERVICE_NAME --task-definition $FAMILY_NAME

