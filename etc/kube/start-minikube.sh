#!/bin/bash

set -Eex

# Parse flags
VERSION=v1.19.0
minikube_args=(
  "--vm-driver=none"
  "--kubernetes-version=${VERSION}"
  "--extra-config=apiserver.enable-admission-plugins=PodSecurityPolicy"
  "--addons=pod-security-policy"
)
while getopts ":v" opt; do
  case "${opt}" in
    v)
      VERSION="v${OPTARG}"
      ;;
    \?)
      echo "Invalid argument: ${opt}"
      exit 1
      ;;
  esac
done

sudo env "PATH=$PATH" "CHANGE_MINIKUBE_NONE_USER=true" \
  minikube start "${minikube_args[@]}"

# Try to connect for three minutes
for _ in $(seq 36); do
  if kubectl version &>/dev/null; then
    exit 0
  fi
  sleep 5
done

# Give up--kubernetes isn't coming up
minikube delete
sleep 30 # Wait for minikube to go completely down
exit 1
