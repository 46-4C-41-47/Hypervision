# Rapport projet annuel
Romain Besson - Tuteur : Martin Lekpa

## Description du projet
Les manières modernes de gérer le déploiement d'un logiciel nécessitent l'utilisation de machines virtuelles pour utiliser au mieux le hardware et les ressources utilisées en fonction du besoin. Cependant cette manière de faire rajoute une couche de complexité et donc requiert des outils de supervisions pour gérer ce parc de machines virtuelles. Ces outils de supervisions sont nombreux, donc pour rendre le travail de ceux qui gèrent le parc plus facile il existe des outils dits d'hypervision dont le but est de regrouper les outils de supervision en un seul outil .
Le but est de créer une solution d'hypervision dans le style de [canopsis](https://www.canopsis.fr/)

## Fonctionnalités
 - Authentification
 Les utilisateurs ont un role qui leur est attribué et qui conditionne les outils auxquels ils ont accès, c'est rôles sont : visiteur, admin et superadmin.
 - Gestion des utilisateurs et des rôles
 Les admins doivent pouvoir voir la liste des utilisateurs. Les superadmins doivent pouvoir voir la liste des utilisateurs , créer, modifier et supprimer des comptes.
 - Système de logs
 On doit avoir un onglet qui liste les évènements se produisent en temps réel sur le système et qui corrèle les informations pour mettre en évidence les premières causes et l'impacte d'un indicent sur le système.
 - Vision de l'état du système 
 On doit avoir une vision synthétique de l'état actuel du système et des ressources utilisées
 - Création de statistiques
 Mise à disposition de statistiques sur l'exécution de calculs ou de processus complexes.

## Choix des technologies
Ce projet utilisera des applications existantes et diverses ce qui peut poser des problèmes d'interfacages entre ces mêmes applications, pour régler ce problème on utilisera Docker ce qui permet d'isoler chaque application dans des conteneurs et de créer et gérer les communications de manière centralisée entre les applications.

Pour le stockage des données on utilisera ElasticSearch qui par son moteur de recherche moderne offre une grande rapidité d'intéraction avec les données.

L'application utilisera un cache pour accélérer l'application, pour ce faire on utilisera Redis qui est une base données en mémoire ce qui lui permet d'utiliser les avantages de la mémoire par rapport au stockage classique à savoir la rapidité.

Pour la partie IHM l'application se présentera dans une page web. Pour créer des pages webs avec Go on utilisera le concept de template qui permet d'interpoler des variables dans un fichier html, pour rendre la page web plus dynamique on utilisera HTMX qui intègre le principe d'Ajax dans du html classique.

Pour l'application principale on utilisera le langage moderne qu'est Go. Les raisons de ce choix sont :
- la simplicité de ce langage qui permet de créer des applications facile à maintenir et à déployer
- la rapidité, même si Go offre des paradigmes de haut niveau comme de la programmation orientée objet il a pour but de créer des applications rapides, ce qui est requit dans le concept de vitualisation qui même si il apporte des avantages apporte aussi le défaut d'un système plus lent par le nombre de couches ajoutées.
- se former à un langage moderne qui est de plus en plus utilisé dans le monde du développement

## Architecture
![Image](./suivie%20de%20projet/Soutenance/Architecture.png "Title")

On peut voir ci-dessus un schéma en pseudo UML qui permet de visualer les différents conteneurs qui composent l'application.

On peut y voir 2 conteneurs qui sont Kibana et Redis Insight qui ne servent qu'au développement pour visualiser et interagir avec les données de leur conteneur associé. 

## Méthodologie de travail
Plusieurs rendez-vous régulier ont été fixés de manière périodique environ une fois par mois. 

Une méthode de travail qui aurait pu être mise en place pour améliorer la qualité et la rapidité de développement et le TDD (Test Driven Devlopment) qui est une méthode qui consiste à écrire les tests avant d'écrire le code d'un composant ce qui permet d'être sûr que le composant répond aux attentes fixées au préalable.

## Fonctionnalités implémentées
L'application est capable de gérer la connexion des utilisateurs, la création et suppression des utilisateurs en fonction de leur rôle.
Il y a 3 rôles : super administrateur, administrateur, et observateur

Upload de données dans la base elastic search
