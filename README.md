# <div style="text-align: center"> Projet HANGMAN</div>


## Sommaire :

* [Description](#Description)
* [La base du projet](#La-base-du-projet)
* [Utilisation](#utilisation)
    * [Installation](#installation)
    * [Lancement](#lancement)
* [Quelques images du jeu](#Quelques-Images-du-jeu)

## Description :

<div style="text-align: justify"> L'objectif de ce projet est de réaliser un jeu du pendu en Golang. Le but est de trouver le mot caché en entrant des lettres (ou le mot entier) avec un nombre de tentatives limité. Si le mot contient la lettre entrée, elle est affichée. Sinon, le joueur perd une vie (ou 2 dans le cas du mot) et le dessin d'un pendu se met à apparaît. Le joueur a 10 vies au total. S'il trouve le mot avant de perdre ses 10 vies, il gagne. Sinon, il perd et le dessin est terminé. </div>

## La base du projet :

<div style="text-align: justify"> Le projet est composé de 3 parties : </div>
<div style="text-align: justify"> - Le fichier main.go qui contient l'initialisation du jeu. </div>
<div style="text-align: justify"> - Le dossier fonctions qui contient les fonctions du jeu ainsi que le fichier word.go </div>
<div style="text-align: justify"> - Le dossier asset qui contient les fichiers textes contenant les mots utilisés par le jeu, l'ASCII Art du hangman, le fichier les lettres en ASCII Art et un fichier save pour pouvoir sauvegarder sa progression. </div>
<div style="text-align: justify"> - Le fichier start.sh qui permet d'avoir un menu de lancement en CLI</div>

## Utilisation :

### *Installation*

<div style="text-align: justify"> Avant de pouvoir lancer une partie, il faut cloner le repository en local avec la commande suivante : </div>

```bash
git clone <url_du_repository>
```
### *Lancement*

<div style="text-align: justify"> Pour lancer une nouvelle partie, il faut se placer dans le dossier du projet et lancer la commande suivante : </div>

```bash
go run main.go asset/words.txt
```


## Fonctionnalités avancées :

<div style="text-align: justify"> Il est possible de sauvegarder une partie afin de la reprendre plus tard. <br>
Pour se faire, il faut se placer dans le dossier du projet et lancer la commande suivante pendant le jeu: </div>

```bash
STOP
```

<div style="text-align: justify"> Pour reprendre sa partie, on se place dans le dossier du projet et on lance la commande suivante : </div>

```bash
go run main.go --startWith asset/save.txt
```
<div style="text-align: justify"> Il est possible de lancer une partie en mode hard, il faut se placer dans le dossier et lancer la commande suivante :</div>

```bash
go run . --hard asset/words.txt
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
