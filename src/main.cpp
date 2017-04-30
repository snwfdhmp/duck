#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include "classes/Emplacement/Emplacement.class.h"
#include "classes/Mouvement/Mouvement.class.h"
#include "classes/Player/Player.class.h"
#include "config/macros.h"
#include "config/constants.h"

using namespace std;

//GRILLE[X][Y]

void clearConsole() {
	int i;
	for (i = 0; i < CONSOLE_HEIGHT/5; ++i)
		printf("\n\n\n\n\n");
}

int printGrille(Emplacement grille[TAILLE][TAILLE]) {
	printf("=========================\n");
	printf("======== PLATEAU ========\n");
	printf("=========================\n");
	int x,y;
	int score[2] = {0, 0};
	printf("=    ==0== ==1== ==2== ==3== ==4== ==5== ==6== ==7== ==8==  =\n");
	for (y = 0;y < TAILLE; ++y)
	{
		printf("= %d: ", y);
		for (x = 0; x < TAILLE; ++x)
		{
			if (grille[x][y].valeur != 2) {
				(grille[x][y].valeur)?printf("\033[31m"):printf("\033[32m");
				printf("(%d:%d) ", grille[x][y].valeur, grille[x][y].hauteur);
				printf("\033[39m");
				score[grille[x][y].valeur]+=grille[x][y].hauteur;
			}
			else
				printf("      ");
		}
		printf(" =\n");
	}
	printf("==== 0 : %d | 1 : %d ====\n", score[0], score[1]);
	return 0;
}

void initGrille(Emplacement grille[TAILLE][TAILLE]) {
	int x,y;

	//on remplit case par case en changeant de couleur à chaque ligne
	for (y = 0; y < TAILLE; ++y)
		for (x = 0; x < TAILLE; ++x)
			grille[x][y].init(y%2);

	//maintenant on définit les cases qui sont vides par défaut
	APPLY_DEFAULT_EMPTY(grille);

}

int onFinish(Emplacement grille[TAILLE][TAILLE], Player* playerA, Player* playerB) {
	printf("OnFinish !\n");
	printGrille(grille);
	//getchar();
	if(playerA->getScore(grille) > playerB->getScore(grille))
		cout << playerA->getName() << " remporte la partie ! (points -> "<< playerA->points++ << "+1)\n";
	else if (playerA->getScore(grille) < playerB->getScore(grille))
		cout << playerB->getName() << " remporte la partie ! (points -> "<< playerB->points++ << "+1)\n";
	else
		printf("Egalité !");

	printf("M:%d | L:%d\n", playerA->points, playerB->points);

	initGrille(grille);
	return 0;
}
int play(int playerAType = PLAYER_TYPE_HUMAN, int playerBType = PLAYER_TYPE_HUMAN) {
	Player playerA, playerB;

	Emplacement grille[TAILLE][TAILLE];

	initGrille(grille);
	//PlayerBot* actual;
	if(ENV_DEV) {
		playerA.init(0, "Martin",playerAType);
		playerB.init(1, "Landry",playerBType);
	} else {
		string nameA, nameB;
		cout << "Player 1 name : ";
		cin >> nameA;
		cout << "Player 2 name : ";
		cin >> nameB;
		playerA.init(0, nameA,playerAType);
		playerB.init(1, nameB,playerBType);
	}

	//TODO : explain C++ tricks
	while(true)
		if(printGrille(grille) && playerA.evaluate(grille) == -1 &&
			printGrille(grille) && playerB.evaluate(grille) == -1)
			onFinish(grille, &playerA, &playerB);
	return 0;
}

int main(int argc, char const *argv[])
{
	int choice;
	printf("Welcome to AVALAM Game !\n");
	printf("Select a game mode :\n");

	do{
	printf("1:Human vs Human\n2:Human vs IA\n3:IA vs IA\n");
	printf("4:Quit\nChoice : ");
	scanf("%d", &choice);
	switch(choice) {
		case 1:
			play(PLAYER_TYPE_HUMAN, PLAYER_TYPE_HUMAN);
			break;
		case 2:
			play(PLAYER_TYPE_HUMAN, PLAYER_TYPE_IA);
			break;
		case 3:
			play(PLAYER_TYPE_IA, PLAYER_TYPE_IA);
			break;
		case 4:
			printf("See you soon !\n");
			break;
		default :
			printf("Please enter a choice between 1 and 4...\n");
			break;
	}
	} while(choice != 4);

	return 0;
}