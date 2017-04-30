#include <stdio.h>
#include "MovePlan.class.h"
#include "../Emplacement/Emplacement.class.h"
#include "../../config/constants.h"

//include macros for unit test
#include "../../config/macros.h"

/*
	UNIT TEST FOR MOVEPLAN
*/


int main(int argc, char const *argv[])
{
	int team;
	Emplacement grille[TAILLE][TAILLE];
	MovePlan a;
	unsigned int err=0;

	grille[1][2].valeur = 1;
	grille[1][3].valeur = 1;
	grille[6][7].valeur = 1;
	grille[7][7].valeur = 1;
	grille[1][1].valeur = 1;
	grille[4][3].valeur = 1;
	grille[4][5].valeur = 1;

	grille[1][2].hauteur = 1;
	grille[1][3].hauteur = 1;
	grille[6][7].hauteur = 1;
	grille[7][7].hauteur = 1;
	grille[1][1].hauteur = 1;
	grille[4][3].hauteur = 1;
	grille[4][5].hauteur = 1;

	for (team = 0; team < 2; ++team)
	{
		SHOULD_BE_FALSE(a.init(-1, -1, -2, -2, 3, grille))
		SHOULD_BE_FALSE(a.init(1, 2, 1, 2, team, grille))

		SHOULD_BE_TRUE(a.init(1, 2, 1, 3, team, grille))
		SHOULD_BE_TRUE(a.init(6, 7, 7, 7, team, grille))
		SHOULD_BE_TRUE(a.init(1, 2, 1, 1, team, grille))
		SHOULD_BE_FALSE(a.init(4, 5, 4, 3, team, grille))
	}

	grille[2][4].valeur = 2;
	grille[2][5].valeur = 2;
	grille[2][6].valeur = 2;
	grille[2][7].valeur = 2;

	grille[2][8].valeur = 1;
	grille[3][7].valeur = 1;

	grille[2][8].hauteur = 2;
	grille[3][7].hauteur = 2;

	for (team = 0; team < 2; ++team)
	{
		SHOULD_BE_FALSE(a.init(2, 3, 2, 4, team, grille))
		SHOULD_BE_FALSE(a.init(3, 4, 2, 4, team, grille))
		SHOULD_BE_FALSE(a.init(2, 4, 2, 5, team, grille))
		SHOULD_BE_FALSE(a.init(2, 4, 2, 3, team, grille))
		SHOULD_BE_FALSE(a.init(2, 5, 2, 5, team, grille))

		SHOULD_BE_TRUE(a.init(1, 2, 1, 3, team, grille))
		SHOULD_BE_TRUE(a.init(2, 8, 3, 7, team, grille))
	}

	UNIT_TEST_RETURN(err)
}