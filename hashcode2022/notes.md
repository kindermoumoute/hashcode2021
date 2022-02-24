Pistes :
- Trier les clients
  - Ratio ingrédients aimés/détestés
  - Pondéré par le nombre d'ingrédients aimés
- Trier les ingrédients
  - Cardinalité aimé/detestés
  - Pondérés par la combinatoire de l'ingrédient
  - Donner des scores pour des couples d'ingrédients (si deux ingrédients sont souvent associés, le couple devient important)
- Utiliser un graphe sur les associations d'ingrédients
- Ajouter/supprimer des ingrédients semi-aléatoirement et repasser sur l'optimisation


Scoring ingrédient :
- Somme des aimés, somme des détestés


Scoring client :
- Pour chaque ingrédient,


Idée d'algo 1 :
1. Ajouter tous les ingrédients non détestés
2. Supprimer tous les ingrédients qui ne sont pas aimés

3. Mettre de côté les clients non résolus dont le ratio d'ingrédients (détesté/aimé > 1)
4. Trier les clients restants (complexité de résolution)
5. Pour chaque client, ajouter les ingrédients et vérifier l'évolution de la cardinalité de clients
6. Pour chaque nouveau client, regarder s'il est éliminé par un ingrédient. Si l'ingrédient est plus souvent détesté, le supprimer.
7. Boucler une fois sur 3.

8. Rajouter les clients mis de côté à 3.
9. Boucler deux fois sur 4.


Idée d'algo 2 :
1. Ajouter tous les ingrédients non détestés
2. Supprimer tous les ingrédients qui ne sont pas aimés
3. Mettre de côté les clients non résolus dont le ratio d'ingrédients (détesté/aimé > 1, ratio général)

4. 
