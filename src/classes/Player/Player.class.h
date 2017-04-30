#ifndef PLAYER_H
#define PLAYER_H

#include <string>

#include "../Emplacement/Emplacement.class.h"

// include global constants
#include "../../config/constants.h"


class Player
{
public:

	int type;
	int team;
	std::string name;
	int delay;
	int points;
	//TODO : int(Player::*evaluate)(Emplacement[TAILLE][TAILLE]) function pointer

	//int pions[TAILLE*TAILLE][2] = {-1};

	int init(int newTeam, std::string newName = "Joueur", int newType = PLAYER_TYPE_MIN); //initialise le joueur (default type -> IA)

	std::string getName(); //returns name

	int getScore(Emplacement grille[TAILLE][TAILLE]); //return le score de l'équipe du joueur

	int evaluate(Emplacement grille[TAILLE][TAILLE]); //input un mouvement par le choisir

	int HumanEvaluate(Emplacement grille[TAILLE][TAILLE]); //evaluate pour le human

	int IAEvaluate(Emplacement grille[TAILLE][TAILLE]); //evaluate pour l'ia
};

#endif