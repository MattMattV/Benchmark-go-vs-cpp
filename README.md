# Benchmark-go-vs-cpp
Le but de ce projet est de voir l'efficacité du Go par raport au C++ sur l'aspect de la programmation concurrente.

### Méthode
Le programme est très simple, on cherche à calculer l'aire d'une fonction en l'approximant avec des rectangles, c'est la méthode de la [somme de Riemann](http://mathworld.wolfram.com/RiemannSum.html). 
On testera successivement avec 1, 2, 4, 8, 16 et 32 cœurs avec un nombre de rectangles croissant.

### Implémentation
L'algorithme a été implanté en Go et en C++ et les programmes utilisent les argument de ligne de commande pour déterminer combien de cœurs utiliser, sur quelle intervalle calculer et le nombre de rectangles à utiliser. Cela facilite grandement l'automatisation des tests.

### Compilation
La gestion des threads utilisée ici étant expérimentale il faudra rajouter les arguments `-std=c++0x -pthread` lors de l'utilisation de `g++`.
