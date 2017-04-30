# Avalam - Projet de 1ère année à Centrale Lille

Ce jeu est notre projet de première année à l'Ecole Centrale de Lille (IG2I).
Nous avons décidé de créer une version jeu vidéo du jeu de société Avalam.
Nous avons également travaillé sur une intelligence artificielle capable de jouer et de battre un joueur humain. Nous essayons toujours de l'optimiser au maximum.
Ce projet est encore en phase de développememnt.

## Pour commencer

Ces instructions vont vous aider à faire tourner le projet sur votre machine dans un but de test. Pour de plus amples informations nous vous invitons à vous rendre dans la section "Analyse du code".

### Pré-requis

Vous aurez besoin de ces programmes déjà installés sur votre machine :

```
g++
Un shell UNIX
```

### Installation

A step by step series of examples that tell you have to get a development env running

Premièrement compilez le projet

```
g++ src/main.cpp src/classes/*/*.class.cpp -o avalam.exe
```

Donnez lui les droits d'éxécution

```
chmod +x avalam.exe
```

Et lancez le jeu !

```
./avalam.exe
```

Voici un aperçu du rendu du jeu en mode console.

```
Running build/release-0.1.2.exe...
=========================
======== BOARD ========
=========================
=    ==0== ==1== ==2== ==3== ==4== ==5== ==6== ==7== ==8==  =
= 0:             (0:1) (0:1)                                =
= 1:       (1:1) (1:1) (1:1) (1:1)                          =
= 2:       (0:1) (0:1) (0:1) (0:1) (0:1) (0:1)              =
= 3:       (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1)  =
= 4: (0:1) (0:1) (0:1) (0:1)       (0:1) (0:1) (0:1) (0:1)  =
= 5: (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1)        =
= 6:             (0:1) (0:1) (0:1) (0:1) (0:1) (0:1)        =
= 7:                         (1:1) (1:1) (1:1) (1:1)        =
= 8:                               (0:1) (0:1)              =
==== 0 : 24 | 1 : 24 ====
Bot John initialised ! (team : 0)
David has joined the game ! (team : 1)
=========================
======== PLATEAU ========
=========================
=    ==0== ==1== ==2== ==3== ==4== ==5== ==6== ==7== ==8==  =
= 0:             (0:1) (0:1)                                =
= 1:       (1:1) (1:1) (1:1) (1:1)                          =
= 2:       (0:1) (0:1) (0:1) (0:1) (0:1) (0:1)              =
= 3:       (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1)  =
= 4: (0:1) (0:1) (0:1) (0:1)       (0:1) (0:1) (0:1) (0:1)  =
= 5: (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1) (1:1)        =
= 6:             (0:1) (0:1) (0:1) (0:1) (0:1) (0:1)        =
= 7:                         (1:1) (1:1) (1:1) (1:1)        =
= 8:                               (0:1) (0:1)              =
==== 0 : 24 | 1 : 24 ====


```

## Built With

* [C++](http://www.cplusplus.com/) - Le Langage
* [SDL2](https://www.libsdl.org/) - Simple DirectMedia Layer 2 (Librairie graphique)

## Contribuer

N'hésitez pas à nous envoyer vos retours sur le code et à nous donner des idées d'améliorations.

## Versioning

Nous utilisons pour le versioning <code>compile</code>, un script personnel. (disponible sur le repo)

## Auteurs

* **Martin Joly** - *Student* - [snwfdhmp](https://github.com/snwfdhmp)

* **Landry Monga** - *Student* - [lvndry](https://github.com/lvndry)

See also the list of [contributors](https://github.com/snwfdhmp/MPI/contributors) who participated in this project.

## License

Ce projet n'est actuellement sous aucune license.

# Explication du projet

## Présentation du repo

Ce repo correspond au projet MPI (projet C de 1ère année à IG2I (Ecole Centrale de Lille)).

Le principe du projet est de coder une version jeu vidéo du jeu de stratégie "Avalam".

## Principe du jeu

Les règles et principes du jeu sont exprimés dans cette vidéo par le créateur du jeu.

[![https://www.youtube.com/watch?v=DbZIvQSyFvA](https://img.youtube.com/vi/DbZIvQSyFvA/0.jpg)](https://www.youtube.com/watch?v=DbZIvQSyFvA)

## Principe de la version jeu vidéo

Les règles du jeu restent les même. Néanmoins nous avons implémenté 3 différents modes de jeu :

- Humain vs Humain :
  2 joueurs humains s'affrontent au tour par tour, en déplaçant les pions à l'aide de la souris
- Humain vs IA
  1 joueur humain contre une intelligence artificielle s'affrontent. L'intelligence artificielle joue de manière automatisée le meilleur mouvement qu'elle peut prédire.
- IA vs IA
  2 IA s'affrontent. L'utilisateur voit de déroulement de la partie et peut l'arrêter, l'accélérer, la ralentir, pour pouvoir observer la manière de jouer de l'IA.
  
## Contraintes

Langages :
    - C
    - C++
   
Librairies :
    - SDL2
    
Compilation :
    - Unix (gcc/g++)
  
 ## Organisation du projet
 
 [![OVERVIEW](https://preview.ibb.co/c9Y7vk/Capture_d_e_cran_2017_04_28_a_19_15_17.png)](https://github.com/snwfdhmp/avalam-ai-game/)
 
Le projet est divisé en plusieurs répertoires :

- [<code>build</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/build)
  Tous nos builds sont sauvegardés dans ce dossier comme tel :
  
  [![OVERVIEW](https://preview.ibb.co/eWkLFk/Capture_d_e_cran_2017_04_28_a_19_29_09.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/build)
  
  Le versioning est géré par l'utilitaire [<code>compile</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/compile) que nous avons créé.
  
- [<code>config</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/config)
  cat-variables pour le fichier <code>compile</code> et peut-être pour d'autres choses plus tard
  
  [![OVERVIEW](https://preview.ibb.co/jsLnvk/Capture_d_e_cran_2017_04_28_a_19_28_53.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/config)
  
- [<code>junk</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/junk)
  Quelques scripts qui nous on servit et qu'on ne supprime pas pour le moment. Ils ne font pas partie du code source
  
  [![OVERVIEW](https://preview.ibb.co/d7e2T5/Capture_d_e_cran_2017_04_28_a_19_29_23.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/junk)
  
- [<code>logs</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/logs)
  Logs de compilation fournis par <code>compile</code> et surement d'autres par la suite.
  
  [![OVERVIEW](https://preview.ibb.co/dF8J1Q/Capture_d_e_cran_2017_04_28_a_19_29_40.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/logs)
  
- [<code>src</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/src)
  Répertoire des sources. Il est organisé comme tel :
  
  [![OVERVIEW](https://preview.ibb.co/fCQLFk/Capture_d_e_cran_2017_04_28_a_19_30_02.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/src)
  
  - [<code>main.cpp</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/src/main.cpp)
    Fichier principal
  - [<code>classes</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/src/classes)
    Contient toutes les classes que nous avons codées pour le projet.
    Chaque classe est représentée par un dossier à son nom contenant au moins
      - <code>*className*.class.cpp</code> => contient les implémentations de fonctions de la classe
      - <code>*className*.class.cpp</code> => contient le prototype de la classe
      - <code>*className*.test.cpp</code> => tests unitaires pour la classe
      
    [![OVERVIEW](https://preview.ibb.co/mth0Fk/Capture_d_e_cran_2017_04_28_a_19_30_14.png)](https://github.com/snwfdhmp/avalam-ai-game/tree/master/src/classes/MovePlan)
    
- [<code>compile</code>](https://github.com/snwfdhmp/avalam-ai-game/tree/master/compile)
  Shell-script utilisé pour la compilation, les logs, le versioning, etc.
