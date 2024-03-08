# Cahier des charges techniques et méthodologiques
## Résumé du projet
Le but du projet est de créer une solution d'hypervision dans le style de [canopsis](https://www.canopsis.fr/), c'est-à-dire un outil qui regroupe tout les informations et outils nécessaires pour gérer un SI (système d'information) comme par exemple des alertes en cas d'incident ou des statistiques sur le SI.

## Technologies utilisées
Ce projet devra se baser sur des conteneurs docker pour pouvoir interfacer les différentes technologies. 
Le cœur de l'application sera un programme en GO qui devra faire l'intermédiaire entre les différents composants. Pour la partie front-end on utilisera une technologie web du nom de HTMX qui repose sur le principe d'AJAX et qui permet d'utiliser un autre langage que JavaScript pour faire du web. 
Pour la partie données on utilisera un conteneur Redis qui permet un accès rapide aux données grâce au fait qu'elles sont stockées en mémoire et pour l'aspect recherche des données on utilisera Elasticsearch, qui est un moteur de recherche dont l'avantage est d'être bien plus rapide que des langages de base de données classique comme MySQL.
L'interface web devra permettre aux utilisateurs de se connecter, de créer et supprimer des utilisateurs pour les super admins ainsi que de voir ces utilisateurs pour les admins.
Elle devra aussi mettre en forme les données stocké dans Redis et restitué par Elasticsearch et d'en faire des statistiques.

## Méthodologie de travail
