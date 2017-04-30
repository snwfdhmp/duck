#ifndef MOUVEMENT_CPP
#define MOUVEMENT_CPP

#include <stdio.h>
#include "../Emplacement/Emplacement.class.h"
#include "Mouvement.class.h"


// include global constants
#include "../../config/constants.h"

int Mouvement::isInTheField(int coor) {
	return (coor < TAILLE && coor >= 0)?0:-1;
}

int Mouvement::isCorrect(int x_s, int y_s, int x_d, int y_d, Emplacement grille[TAILLE][TAILLE]) {
	if(isInTheField(x_s) == -1 || isInTheField(y_s) == -1 || isInTheField(x_d) == -1 || isInTheField(y_d) == -1)
		return -1;
	if(grille[x_s][y_s].valeur == 2 || grille[x_d][y_d].valeur == 2)
		return -1;
	 if (!(x_s-x_d >= -1 && x_s-x_d <= 1 && y_s-y_d >= -1 && y_s-y_d <= 1))
	 	return -1;
	    //if(!field)
	    //	printf("==!Position incorrecte [%d:%d]=>[%d:%d]!==\n", x_s, y_s, x_d, y_d);
	if(x_s == x_d && y_s == y_d)
		return -1;
		//if(!different)
		//	printf("==!C'est la même pièce!==\n");
	//printf("%d, %d, %d\n", grille[x_d][y_d].hauteur, grille[x_s][y_s].hauteur, grille[x_s][y_s].hauteur+grille[x_d][y_d].hauteur);
	if(!(grille[x_d][y_d].hauteur > 0 && grille[x_s][y_s].hauteur > 0 && grille[x_s][y_s].hauteur+grille[x_d][y_d].hauteur <= 5))
		return -1;
		//if(!hauteurLegale && grille[x_d][y_d].hauteur > 0 && grille[x_s][y_s].hauteur > 0)
		//	printf("grille[x_d][y_d].hauteur [%d] > 0 && grille[x_s][y_s].hauteur [%d] > 0 && grille[x_s][y_s].hauteur+grille[x_d][y_d].hauteur [%d] <= 5\n",grille[x_d][y_d].hauteur, grille[x_s][y_s].hauteur, grille[x_s][y_s].hauteur+grille[x_d][y_d].hauteur);
	//printf("OK\n");
	return 0;
}

int Mouvement::init(int x_s, int y_s, int x_d, int y_d, Emplacement grille[TAILLE][TAILLE]) { //s = src, d= des 
	//printf("\tinit %d %d %d %d\n", x_s, y_s, x_d, y_d);
	if (isCorrect(x_s, y_s, x_d, y_d, grille) == -1)
		return -1;

	src[0] = x_s;
	src[1] = y_s;
	des[0] = x_d;
	des[1] = y_d;

	return 0;
}

int Mouvement::apply(Emplacement grille[TAILLE][TAILLE]) {
	if(isCorrect(src[0], src[1], des[0], des[1], grille) == -1)
		return -1;
		//printf("grille[%d][%d].valeur (:%d) = grille[%d][%d].valeur (:%d)\n", des[0], des[1], grille[des[0]][des[1]].valeur, src[0], src[1], grille[src[0]][src[1]].valeur);
		//printf("grille[%d][%d].hauteur (:%d) += grille[%d[%d].hauteur (:%d)\n", des[0], des[1], grille[des[0]][des[1]].hauteur, src[0], src[1], grille[src[0]][src[1]].hauteur);
	grille[des[0]][des[1]].valeur = grille[src[0]][src[1]].valeur;
	grille[des[0]][des[1]].hauteur += grille[src[0]][src[1]].hauteur;
	grille[src[0]][src[1]].valeur = 2;
	grille[src[0]][src[1]].hauteur = 0;
		//printf("executing\n");
	return 0;
}

int Mouvement::verify(Emplacement grille[TAILLE][TAILLE]) {
	return init(src[0], src[1], des[0], des[1], grille);
}


#endif