# Comment exécuter le projet

Cloner ce repos, puis une fois fais placez vous dans le répertoire **HYPERVISION**
puis exécutez les commandes suivantes depuis un terminal (il vous faudra avoir installé [docker](https://www.docker.com/products/docker-desktop/) au préalable)
``` docker build -t gowebappimage . ```
pour construire l'image go.

Puis exécutez la commande ``` docker compose up ```.

Ou alors vous pouvez exécuter la commande ./hypervision.bat pour lancer un script qui fait ça à votre place.
Vous n'avez ensuite plus qu'à vous rendre sur votre [localhost](http://localhost:3000) pour accéder à l'application web
