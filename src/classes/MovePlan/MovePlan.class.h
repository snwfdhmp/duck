#ifndef MOVEPLAN_H
#define MOVEPLAN_H


#include <time.h>
#include <stdlib.h>
#include "../Mouvement/Mouvement.class.h"
#include "../Emplacement/Emplacement.class.h"

// include global constants
#include "../../config/constants.h"

class MovePlan
{
public:
	Mouvement mvt;
	int score, team;

	MovePlan(); //constructeur

	//Initialise le MovePlan
	int init(int x_s, int y_s, int x_d, int y_d, int newTeam, Emplacement grille[TAILLE][TAILLE]);

	//Calcule les points d'une équipe sur une grille
	int getPoints(Emplacement grille[TAILLE][TAILLE], int whichTeam);

	//Calcule le score du MovePlan
	int calcScore(Emplacement grille[TAILLE][TAILLE]);

	//Renvoie le score du MovePlan
	int getScore();
};

#endif