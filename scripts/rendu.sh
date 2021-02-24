#!/bin/bash

go run . input/c_many_ingredients.in


zip -r ./output/rendu-$(date +%H%M).zip $(find . -maxdepth 1 -not -name "." -not -name "input" -not -name ".idea" -not -name ".DS_Store" -not -name "go.mod" -not -name "go.sum" -not -name ".gitignore" -not -name "scripts" -not -name "trash" -not -name ".git" -not -name "output")
