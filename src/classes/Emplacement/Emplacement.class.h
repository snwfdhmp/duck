#ifndef EMPLACEMENT_H
#define EMPLACEMENT_H

class Emplacement
{
public:
	int valeur; // 0 ou 1
	int hauteur; // entre 0 et 5

	void init(int val); // Initialise l'emplacement avec sa valeur

	bool isVide(); // renvoie si l'emplacement est vide
	
};

#endif