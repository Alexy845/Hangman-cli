# <div style="text-align: center"> Projet HANGMAN</div>


## Sommaire :

* [Description](#Description)
* [La base du projet](#La-base-du-projet)
* [Utilisation](#utilisation)
* [Quelques images du jeu](#Quelques-Images-du-jeu)

## Description

<div style="text-align: justify"> Ce projet est un jeu du pendu en Golang. Le but est de trouver le mot caché en entrant des lettres (ou le mot entier). Si le mot contient la lettre entrée, elle est affichée. Sinon, le joueur perd une vie (ou 2 dans le cas du mot). Le joueur a 10 vies au total. Si le joueur trouve le mot avant de perdre ses 10 vies, il gagne. Sinon, il perd. </div>

## La base du projet

<div style="text-align: justify"> Le projet est composé de 3 parties : </div>
<div style="text-align: justify"> - Le fichier main.go qui contient l'initialisation du jeu. </div>
<div style="text-align: justify"> - Le dossier fonctions qui contient les fonctions du jeu. </div>
<div style="text-align: justify"> - Le dossier asset qui contient les fichiers textes contenant les mots utilisés par le jeu, l'ASCII Art du hangman, le fichier les lettres en ASCII Art et un fichier save pour pouvoir sauvegarder sa progression. </div>

## Utilisation

<div style="text-align: justify"> Pour lancer une nouvelle partie, il faut se placer dans le dossier du projet et lancer la commande suivante : </div>

```bash
go run main.go asset/words.txt
```

<div style="text-align: justify"> Pour sauvegarder une partie, il faut se placer dans le dossier du projet et lancer la commande suivante pendant le jeu: </div>

```bash
STOP
```

<div style="text-align: justify"> Pour reprendre sa partie, il faut se placer dans le dossier du projet et lancer la commande suivante : </div>

```bash
go run main.go --startWith asset/save.txt
```

## Quelques Images du jeu
![damn](images/damn-sarcasm.gif)
<div style="text-align: justify"> Voici quelques images du jeu : </div>

### <div style="text-align: justify"> - Lancement du jeu : </div>

![Lancement du jeu](images/debut_jeu.png)

### <div style="text-align: justify"> - Perdre une vie </div>

![Perdre une vie](images/une_vie_perdue.png)

### <div style="text-align: justify"> - Gagner la partie </div>

![Gagner la partie](images/win.png)
### <div style="text-align: justify"> - Perdre la partie </div>

![Perdre la partie](images/lose.png)
