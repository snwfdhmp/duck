#ifndef GRAPHICCOMPONENT_H
#define GRAPHICCOMPONENT_H

class GraphicComponent
{
public:
	int x, y;
	int height, width;
	
	GraphicComponent(int set_x, int set_y, int set_width, int set_height);

	~GraphicComponent();
	
	void onClick(); // click handler function

	void onMouseOver(); // mouseover handler function

	void onMouseOut(); //mouseout handler function
	
};
#endif