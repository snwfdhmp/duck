#ifndef DISPLAY_CPP
#define DISPLAY_CPP

#include "../GraphicComponent/GraphicComponent.class.h"

class Display
{
public:
	GraphicComponent** components;
	int x, y, width, height;
	unsigned int size;

	Display(int set_x, int set_y, int set_width, int set_height);
	
	~Display();

	Display* initWindow();

	Display* updateWindow();
	
	GraphicComponent* addComponent(GraphicComponent* componentToAdd);

	GraphicComponent* getTargeted(int mouse_x, int mouse_y);

};
#endif