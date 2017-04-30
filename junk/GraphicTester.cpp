#include <stdio.h>
#include <stdlib.h>
#include "Display.class.h"
#include "GraphicComponent.cpp"

// TODO -> unit test for this class

int main(int argc, char const *argv[])
{
	Display* view = new Display(0, 0, 1080, 780); // Display(int set_x, int set_y, int set_width, int set_height) 
	GraphicComponent* button = new GraphicComponent(50, 50, 240, 160);
	view->addComponent(button);

	int input_x = 0, input_y = 0;

	while(true) {
		printf("Please enter coordinates for mouse click :\nx:");
		scanf("%d", &input_x);
		printf("y:");
		scanf("%d", &input_y);
		GraphicComponent* aimed = view->getTargeted(input_x, input_y);
		if(aimed == NULL)
			printf("Rien ne se passe...\n");
		else
			aimed->onClick();
	}
	return 0;
}