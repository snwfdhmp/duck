#ifndef PLAYER_CPP
#define PLAYER_CPP

#include <stdio.h>
#include <iostream>
#include <string>
#include "../Emplacement/Emplacement.class.h"
#include "../Mouvement/Mouvement.class.h"
#include "../MovePlan/MovePlan.class.h"
#include "Player.class.h"

// include global constants
#include "../../config/constants.h"
#include "../../config/macros.h"

using namespace std;

int Player::init(int newTeam, string newName, int newType) {
	if((!AVAILABLE_TEAM(newTeam)) || newName.empty() || (!AVAILABLE_TYPE(newType)))
		return -1;
	team = newTeam;
	name = newName;
	type = newType;
	/* TODO : int(Player::*evaluate)(Emplacement[TAILLE][TAILLE]) function pointer
	switch(type) {
		case PLAYER_TYPE_HUMAN:
			evaluate=&Player::HumanEvaluate;
		break;
		case PLAYER_TYPE_IA:
			evaluate=&Player::IAEvaluate;
		break;
		default :
			evaluate=&Player::HumanEvaluate;
		break;
	}*/
	points = 0;
	//cout << name << " a rejoint la partie ! (équipe : "<< team <<")\n";
	return 0;
}

string Player::getName() {
	return name;
}

int Player::getScore(Emplacement grille[TAILLE][TAILLE]) {
	int score=0, x, y;
	for (x = 0; x < TAILLE; ++x)
		for (y = 0; y < TAILLE; ++y)
			if(grille[x][y].valeur == team) score++;
		return score;
}

int Player::HumanEvaluate(Emplacement grille[TAILLE][TAILLE]) {
	Mouvement mvt;
	int x_s, y_s, x_d, y_d;

	puts("====Quel pion bouger?====\n");
	puts("> x : ");
	scanf("%d", &x_s);
	puts("> y : ");
	scanf("%d", &y_s);
	puts("==Vers quelle position?==\n");
	puts("> x : ");
	scanf("%d", &x_d);
	puts("> y : ");
	scanf("%d", &y_d);

	if(mvt.init(x_s, y_s, x_d, y_d, grille) == -1)
		return -1;

	mvt.apply(grille);
	return 0;
}

int Player::evaluate(Emplacement grille[TAILLE][TAILLE]) {
	switch(type) {
		case PLAYER_TYPE_HUMAN:
			return HumanEvaluate(grille);
		case PLAYER_TYPE_IA:
			return IAEvaluate(grille);
		default:
			return HumanEvaluate(grille);
	}
}

int Player::IAEvaluate(Emplacement grille[TAILLE][TAILLE]) {
	int x, y, k, l, len=0;
	MovePlan bestmove, tmp;
	int big = 0;

	for (x = 0; x < TAILLE; ++x)
		for (y = 0; y < TAILLE; ++y)
			if(grille[x][y].valeur == 2)
				continue;
			else
				for (k = -1; k <= 1; ++k)
					for (l = -1; l <= 1; ++l)
						if(grille[x+l][y+k].valeur == 2) continue;
					else
						if(tmp.init(x, y, x+l, y+k, team, grille) > bestmove.getScore())
							bestmove.init(x, y, x+l, y+k, team, grille);

						bestmove.calcScore(grille);
						if(bestmove.score == -1) return -1;
						bestmove.mvt.apply(grille);
						return 0;
					}

#endif