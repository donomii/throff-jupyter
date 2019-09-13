#!/bin/sh

go install .
mkdir -p ~/Library/Jupyter/kernels/throff-jupyter
cp kernel/* ~/Library/Jupyter/kernels/throff-jupyter
