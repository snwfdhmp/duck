#ifndef EMPLACEMENT_CPP
#define EMPLACEMENT_CPP

#include <stdio.h>
#include <stdlib.h>
#include "Emplacement.class.h"

#define TAILLE 9

void Emplacement::init(int val) {

if(val == 2) //si on souhaite créer une case vide
	hauteur = 0;
else { //TODO intégrer vérification de valeur
	if(val != 0 && val != 1)
		return;
	valeur = val;
	hauteur = 1;
}

}

bool Emplacement::isVide() {
	return hauteur==0;
}



#endif