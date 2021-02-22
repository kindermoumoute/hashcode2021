#!/bin/bash

zip -r ./rendu/rendu-$(date +%H%M).zip $(find . -depth 1 -not -name "inputs" -not -name ".git" -not -name "rendu" -not -name ".idea" -not -name "outputs")

