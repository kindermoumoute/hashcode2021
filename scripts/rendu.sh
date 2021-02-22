#!/bin/bash

go run . input/b_little_bit_of_everything.in input/a_example.in input/c_many_ingredients.in input/d_many_pizzas.in input/e_many_teams.in


zip -r ./output/rendu-$(date +%H%M).zip $(find . -maxdepth 1 -not -name "." -not -name "input" -not -name ".idea" -not -name ".DS_Store" -not -name "go.mod" -not -name "go.sum" -not -name ".gitignore" -not -name "scripts" -not -name "trash" -not -name ".git" -not -name "output")
