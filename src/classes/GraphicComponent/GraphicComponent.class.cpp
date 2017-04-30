#ifndef GRAPHICCOMPONENT_CPP
#define GRAPHICCOMPONENT_CPP

#include <stdio.h>
#include "GraphicComponent.class.h"

GraphicComponent::GraphicComponent(int set_x, int set_y, int set_width, int set_height) {
	x = set_x;
	y = set_y;
	width = set_width;
	height = set_height;
}
GraphicComponent::~GraphicComponent(){};

void GraphicComponent::onClick() {
	printf("GraphicComponent at [%d;%d] has fired the onClick() function.\n", x, y);
}

void GraphicComponent::onMouseOver() {
	printf("GraphicComponent at [%d;%d] has fired the onMouseOver() function.\n", x, y);
}

void GraphicComponent::onMouseOut() {
	printf("GraphicComponent at [%d;%d] has fired the onMouseOut() function.\n", x, y);
}

#endif