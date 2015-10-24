#!/bin/bash

cat > cities-controller.json <<EOF
{
  "kind": "ReplicationController",
  "apiVersion": "v1",
  "metadata": {
    "name": "cities",
    "labels": {
      "name": "cities"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "name": "cities"
    },
    "template": {
      "metadata": {
        "labels": {
          "name": "cities",
          "deployment": "${WERCKER_GIT_COMMIT}"
        }
      },
      "spec": {
        "containers": [
          {
            "imagePullPolicy": "Always",
            "image": "quay.io/grandmore/wercker:${WERCKER_GIT_COMMIT}",
            "name": "cities",
            "ports": [
              {
                "name": "http-server",
                "containerPort": 5001,
                "protocol": "TCP"
              }
            ]
          }
        ]
      }
    }
  }
}
EOF
