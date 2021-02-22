#!/bin/bash

zip -r ./output/rendu-$(date +%H%M).zip $(find . -depth 1 -not -name "input" -not -name ".idea" -not -name ".DS_Store" -not -name ".gitignore" -not -name "scripts" -not -name ".git" -not -name "output")
