#ifndef PLAYER_IA_TEST_CPP
#define PLAYER_IA_TEST_CPP

#include <stdio.h>
#include "Player.class.h"

//include macro for testing
#include "../../config/macros.h"

int main(int argc, char const *argv[])
{
	Player a;
	unsigned char team;
	int err=0;

	for (team = 0; team < 2; ++team)
	{
		SHOULD_BE_TRUE(a.init(team, "Martin", PLAYER_TYPE_HUMAN))
		SHOULD_BE_TRUE(a.init(team, "L", PLAYER_TYPE_HUMAN))
		SHOULD_BE_TRUE(a.init(team, "MARTIN", PLAYER_TYPE_HUMAN))
		SHOULD_BE_FALSE(a.init(team, "", PLAYER_TYPE_HUMAN))
		SHOULD_BE_FALSE(a.init(team, "", PLAYER_TYPE_IA))
		SHOULD_BE_TRUE(a.init(team, " ", PLAYER_TYPE_IA))
		SHOULD_BE_TRUE(a.init(team, "Test", PLAYER_TYPE_IA))
		SHOULD_BE_TRUE(a.init(team, " ", PLAYER_TYPE_Human))
		SHOULD_BE_FALSE(a.init(team, "Martin", PLAYER_TYPE_MAX+1))
		SHOULD_BE_FALSE(a.init(team, "", PLAYER_TYPE_MAX+1))
		SHOULD_BE_FALSE(a.init(team, " ", PLAYER_TYPE_MAX+2))
	}

	SHOULD_BE_FALSE(a.init(3, "Martin", PLAYER_TYPE_HUMAN))

	UNIT_TEST_RETURN(err)
}

#endif