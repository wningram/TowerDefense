# Vocabulary
- Player: (* or + or =|) a dynamic/moveable object on the Battlefield
- Battlefield: The drawn environment of the game
- Enemy: (*) a player that moves rightward
- Ally: (+) a player that moves leftward
- UserPlayer: (=|) a player that dispenses allies and cannot be hit by an enemy player or the game will end
# Bugs
1. If a enemy hits the wall and an ally moves through teh deactivated enemy, the wall is reset
    - FIXED
2. When an enemy and ally collide, the enemy overwrites the ally and continues on its path
    - FIXED
3. When enemy collides with inactive enemy on teh defense line, the original enemy dissappears into the defense line
    - FIXED