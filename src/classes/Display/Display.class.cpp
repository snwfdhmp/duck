#ifndef DISPLAY_H
#define DISPLAY_H

#include <stdio.h>
#include <stdlib.h>

#include "../GraphicComponent/GraphicComponent.class.h"
#include "Display.class.h"

// include global constants
#include "../../config/constants.h"

Display::Display(int set_x, int set_y, int set_width, int set_height) {
	components = (GraphicComponent**) malloc(sizeof(GraphicComponent*));
	size = 0;
	x = set_x;
	y = set_y;
	width = set_width;
	height = set_height;

};
Display::~Display() {};

Display* Display::initWindow() {
	bool error = false;

		/*
			 SDL initialisation de la fenêtre

			 Mettre error = true s'il y a eu une erreur
			 				(qui empeche de continuer)

		*/

	if(error)
		return NULL;
	return this;
};

Display* Display::updateWindow() {
	bool error = false;
		// SDL clear la fenetre

	for (int i = 0; i < size; ++i)
	{
			// Ajouter le components[i] à la fenetre
			// error = true si erreur
	}

	if(error)
		return NULL;
	return this;
};

GraphicComponent* Display::addComponent(GraphicComponent* componentToAdd) {
	components = (GraphicComponent**) realloc(components, sizeof(GraphicComponent*) * size + 1);
	components[size] = componentToAdd;
	printf("Component [%d;%d] has been added to display [%d;%d] (%d:%d) [size : %d]\n", components[size]->x, components[size]->y, x, y, width, height, size);
	return components[size++]; //size is incremented after he gets injected as index of array (post inc)
};

GraphicComponent* Display::getTargeted(int mouse_x, int mouse_y) {
	if(mouse_x < x || mouse_x > x+width || mouse_y < y || mouse_y > y+height)
		return NULL;
	else
		for (int i = 0; i < size; ++i)
			if(mouse_x < components[i]->x || mouse_x > components[i]->x+components[i]->width || mouse_y < components[i]->y || mouse_y > components[i]->y+components[i]->height)
				continue;
			else
				return components[i];
	return NULL;
};

#endif