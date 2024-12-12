#!/bin/bash

# VARIABLES
RESOURCE_PATH="./.."

echo "################################"
echo "## Create Virtual Environment ##"
echo "################################"
if [ -d "$RESOURCE_PATH/venv" ]; then
  echo "virtual environment 'venv' already exists"
else
  virtualenv $RESOURCE_PATH/venv
fi

echo "################################"
echo "## Install Dependencies       ##"
echo "################################"
if [ -f "./requirements.txt" ]; then
  $RESOURCE_PATH/venv/bin/pip install -r requirements.txt
else
  echo "no requirements.txt file can be found"
fi