#!/bin/bash

if [[ "$OSTYPE" == "darwin"* ]]; then
  zip -r ./output/rendu-$(date +%H%M).zip $(find . -depth 1 -name "*go*")
else
  zip -r ./output/rendu-$(date +%H%M).zip $(find . -maxdepth 1 -name "*go*")
fi